package cloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/util"
	"github.com/disintegration/imaging"
	"google.golang.org/api/option"
)

type clientUploader struct {
	cl         *storage.Client
	log        logger.Logger
	projectID  string
	bucketName string
}

func NewCloudClient(cfg *config.Config, log logger.Logger) Service {
	decoded, _ := base64.StdEncoding.DecodeString(cfg.GoogleCloudServiceAccount)
	client, _ := storage.NewClient(context.Background(), option.WithCredentialsJSON(decoded))
	return &clientUploader{
		cl:         client,
		projectID:  cfg.GoogleCloudProjectID,
		bucketName: cfg.GoogleCloudBucketName,
	}
}

func (c *clientUploader) HostImageToGCS(imageUrl string, name string) (string, error) {
	// check if image is .webp
	if !imageUrlCheck(imageUrl) {
		return imageUrl, nil
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// download image
	err := util.DownloadFile(imageUrl, "temp")
	if err != nil {
		c.log.Errorf(err, "[cloud.HostImageToGCS] failed to download image: %s", err)
		return "", fmt.Errorf("[cloud.HostImageToGCS] failed to download image: %s", err)
	}
	defer os.Remove("temp")
	defer os.Remove("resized.png")

	// get cloud storage bucket handler
	handler := c.cl.Bucket(c.bucketName).Object(fmt.Sprintf("%s.png", name))
	if handler == nil {
		c.log.Errorf(err, "[cloud.HostImageToGCS] failed to find bucket %s: %s", c.bucketName, err)
		return "", fmt.Errorf("[cloud.HostImageToGCS] failed to find storage bucket: %s", err)
	}

	// resize image if from google
	if strings.Contains(imageUrl, "googleusercontent") {
		src, err := imaging.Open("temp")
		if err != nil {
			c.log.Errorf(err, "[cloud.HostImageToGCS] failed to resize image: %s", err)
			return "", fmt.Errorf("[cloud.HostImageToGCS] failed to resize image: %s", err)
		}
		src = imaging.Resize(src, 300, 0, imaging.Lanczos)
		_ = imaging.Save(src, "resized.png")
	}

	// open downloaded image
	file, err := os.Open("resized.png")
	if err != nil {
		c.log.Errorf(err, "[cloud.HostImageToGCS] failed to open image: %s", err)
		return "", fmt.Errorf("[cloud.HostImageToGCS] failed to open image: %s", err)
	}
	defer file.Close()

	// copy downloaded image to storage
	wc := handler.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		c.log.Errorf(err, "[cloud.HostImageToGCS] failed to write to bucket %s: %s", c.bucketName, err)
		return "", fmt.Errorf("[cloud.HostImageToGCS] failed to write file to storage: %s", err)
	}
	defer wc.Close()

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s.png", c.bucketName, name), nil
}

func imageUrlCheck(imageUrl string) bool {
	if strings.Contains(imageUrl, "googleusercontent") {
		return true
	}
	if strings.Contains(imageUrl, ".webp") {
		return true
	}
	return false
}
