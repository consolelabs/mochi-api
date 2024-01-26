package entities

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/consolelabs/mochi-typeset/typeset"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/kafka/message"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	productbotcommand "github.com/defipod/mochi/pkg/repo/product_bot_command"
	productchangelogs "github.com/defipod/mochi/pkg/repo/product_changelogs"
	productchangelogsview "github.com/defipod/mochi/pkg/repo/product_changelogs_view"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) ProductBotCommand(req request.ProductBotCommandRequest) ([]model.ProductBotCommand, error) {
	return e.repo.ProductBotCommand.List(productbotcommand.ListQuery{
		Code:  req.Code,
		Scope: req.Scope,
	})
}

func (e *Entity) ProductChangelogs(req request.ProductChangelogsRequest) ([]model.ProductChangelogs, int64, error) {
	return e.repo.ProductChangelogs.List(productchangelogs.ListQuery{
		Product: req.Product,
		Size:    int(req.Size),
		Page:    int(req.Page),
	})
}

func (e *Entity) GetProductChangelogByVersion(version string) (*model.ProductChangelogs, error) {
	changelog, err := e.repo.ProductChangelogs.GetByVersion(version)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		e.log.Error(err, "[entity.GetProductChangelogByVersion] - repo.ProductChangelogs.GetByVersion() failed")
		return nil, err
	}

	nextVer, _ := e.repo.ProductChangelogs.GetNextVersion(changelog.Id)
	prevVer, _ := e.repo.ProductChangelogs.GetPreviousVersion(changelog.Id)
	changelog.NextVersion = nextVer
	changelog.PreviousVersion = prevVer

	return changelog, nil
}

func (e *Entity) CreateProductChangelogsView(req request.CreateProductChangelogsViewRequest) (*model.ProductChangelogView, error) {
	productchangelogsview := &model.ProductChangelogView{
		Key:           req.Key,
		ChangelogName: req.ChangelogName,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	return productchangelogsview, e.repo.ProductChangelogsView.Create(productchangelogsview)
}

func (e *Entity) GetProductChangelogsView(req request.GetProductChangelogsViewRequest) ([]model.ProductChangelogView, error) {
	return e.repo.ProductChangelogsView.List(productchangelogsview.ListQuery{
		Key:           req.Key,
		ChangelogName: req.ChangelogName,
	})
}

func (e *Entity) CrawlChangelogs() {
	repos := e.svc.Github.GetContents()
	err := e.repo.ProductChangelogs.DeleteAll()
	if err != nil {
		e.log.Error(err, "[entity.CrawlChangelogs()] - cannot delete all")
		return
	}
	// 1. crawl changelog
	for _, repo := range repos {
		// 1.1 validate file markdown
		if !util.ValidateFileMarkdown(repo.Name) {
			continue
		}

		// 1.2 get detail content of repo
		repoDetail, err := e.svc.Github.GetContentByPath(repo.URL)
		if err != nil || repoDetail == nil {
			e.log.Fields(logger.Fields{"title": repo.Name}).Error(err, "[entity.CrawlChangelogs()] - cannot get content of repo")
			continue
		}

		// 1.3 parse content from base64 to string
		rawDecodedText, err := base64.StdEncoding.DecodeString(repoDetail.Content)
		if err != nil {
			e.log.Fields(logger.Fields{"title": repo.Name}).Error(err, "[entity.CrawlChangelogs()] - cannot decode content of repo")
			continue
		}

		// 1.4 convert content string to model.ProductChangelogs
		changelogs := e.parseChangelogsContent(string(rawDecodedText))
		if changelogs == nil {
			e.log.Fields(logger.Fields{"title": repo.Name}).Error(err, "[entity.CrawlChangelogs()] - cannot parse content of repo")
			continue
		}
		changelogs.GithubUrl = repo.HTMLURL
		changelogs.FileName = repo.Name
		changelogs.IsExpired = false

		// 1.5 store changelogs
		err = e.repo.ProductChangelogs.Create(changelogs)
		if err != nil {
			e.log.Fields(logger.Fields{"title": repo.Name}).Error(err, "[entity.CrawlChangelogs()] - cannot store repo")
			continue
		}
	}

	// 2. find new changelog
	newChangelogs, err := e.repo.ProductChangelogs.GetNewChangelog()
	if err != nil {
		e.log.Error(err, "[entity.CrawlChangelogs()] - cannot find new changelog")
		return
	}

	var productChangelogSnapshots []model.ProductChangelogSnapshot
	//var newChangelogMessages []message.NewChangelog
	for _, pc := range newChangelogs {
		productChangelogSnapshots = append(productChangelogSnapshots, model.ProductChangelogSnapshot{
			Filename:  pc.FileName,
			IsPublic:  false,
			CreatedAt: pc.CreatedAt,
			UpdatedAt: pc.UpdatedAt,
		})
	}

	// 3. update product changelog snapshot
	if len(productChangelogSnapshots) > 0 {
		err = e.repo.ProductChangelogs.InsertBulkProductChangelogSnapshot(productChangelogSnapshots)
		if err != nil {
			e.log.Error(err, "[entity.CrawlChangelogs()] - cannot insert bulk product changelog snapshot")
			return
		}
	}

	// 4. get changelog not confirmed
	changelogNotConfirmed, err := e.repo.ProductChangelogs.GetChangelogNotConfirmed()
	if err != nil {
		e.log.Error(err, "[entity.CrawlChangelogs()] - cannot find changelog has not confirmed")
		return
	}

	for _, pc := range changelogNotConfirmed {
		content, images := e.ParseChangelogContent(pc.Version, pc.Title, pc.Content)
		err := e.SendChangelogToChannel(pc.FileName, pc.Version, content, images)

		if err != nil {
			e.log.Error(err, "[entity.CrawlChangelogs()] - cannot send changelog to channel")
			continue
		}
	}

}

func (e *Entity) parseChangelogsContent(content string) *model.ProductChangelogs {
	var changelogs model.ProductChangelogs
	contentSplit := strings.Split(content, "---")
	if len(contentSplit) < 3 {
		return nil
	}

	// 1. get title, date, product, thumbnail url in contentSplit[1]
	table := strings.TrimSpace(contentSplit[1])
	rows := strings.Split(table, "\n")
	for _, row := range rows {
		cRow := strings.Split(row, ": ")
		if len(cRow) < 2 {
			continue
		}

		switch cRow[0] {
		case "date":
			dateString := strings.TrimSpace(cRow[1])
			date, err := time.Parse("02-01-2006", dateString)
			if err != nil {
				changelogs.CreatedAt = time.Now()
				changelogs.UpdatedAt = time.Now()
			}
			changelogs.CreatedAt = date
			changelogs.UpdatedAt = date
		case "title":
			changelogs.Title = strings.TrimSpace(cRow[1])
		case "product":
			changelogs.Product = strings.TrimSpace(cRow[1])
		case "thumbnail_url":
			changelogs.ThumbnailUrl = strings.TrimSpace(cRow[1])
		case "field_version":
			changelogs.Version = strings.TrimSpace(cRow[1])
		case "seo_description":
			changelogs.SeoDescription = strings.TrimSpace(cRow[1])
		}
	}

	// 2. Get content changelogs
	changelogs.Content = strings.TrimSpace(contentSplit[2])

	return &changelogs
}

func (e *Entity) GetProductHashtag(req request.GetProductHashtagRequest) (*model.ProductHashtagAlias, error) {
	data, err := e.repo.ProductHashtag.GetByAlias(req.Alias)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &model.ProductHashtagAlias{
		ProductHashtag:   data,
		ProductHashtagId: data.Id,
		Alias:            strings.ToLower(req.Alias),
	}, nil
}

func (e *Entity) GetProductTheme(req request.GetProductThemeRequest) ([]model.ProductTheme, error) {
	return e.repo.ProductTheme.Get()
}

func (e *Entity) PublishChangeLog(req request.ProductChangelogSnapshotRequest) error {
	changelog, err := e.repo.ProductChangelogs.GetChangelogByFilename(req.ChangelogName)
	if err != nil {
		e.log.Errorf(err, "[entity.PublishChangeLog] - e.repo.ProductChangelogs.GetChangelogByFilename failed")
		return err
	}

	changelogSnapshot, err := e.repo.ProductChangelogs.GetChangelogSnapshotByFilename(req.ChangelogName)
	if err != nil {
		e.log.Errorf(err, "[entity.PublishChangeLog] - e.repo.ProductChangelogs.GetChangelogSnapshotByFilename failed")
		return err
	}

	if !changelogSnapshot.IsPublic {
		msg := message.NewChangelog{
			Type: typeset.NOTIFICATION_NEW_CHANGELOG,
			NewChangelogMetadata: message.NewChangelogMetadata{
				Product:      changelog.Product,
				Title:        changelog.Title,
				Content:      changelog.Content,
				FileName:     changelog.FileName,
				GithubUrl:    changelog.GithubUrl,
				Version:      changelog.Version,
				ThumbnailUrl: changelog.ThumbnailUrl,
				IsExpired:    changelog.IsExpired,
				CreatedAt:    changelog.CreatedAt,
				UpdatedAt:    changelog.UpdatedAt,
			},
		}

		// 3. push notification for new changelogs
		// TODO. implement push notification
		byteNotification, _ := json.Marshal(msg)

		err = e.kafka.ProduceNotification(e.cfg.Kafka.NotificationTopic, byteNotification)
		if err != nil {
			e.log.Errorf(err, "[entity.PublishChangeLog] - e.kafka.Produce failed")
			return err
		}

	}

	err = e.repo.ProductChangelogs.UpdateProductChangelogSnapshot(productchangelogs.ProductChangelogSnapshotQuery{
		Filename: req.ChangelogName,
		IsPublic: req.IsPublic,
	})

	if err != nil {
		e.log.Errorf(err, "[entity.PublishChangeLog] - e.repo.ProductChangelogs.UpdateProductChangelogSnapshot failed")
		return err
	}

	return nil
}

func (e *Entity) SendChangelogToChannel(filename string, version string, content string, images []string) error {
	var image string
	var mochiAvatar string

	if len(images) > 0 {
		image = images[0]
	}

	emojis, err := e.GetEmojiByCode("MOCHI_CIRCLE")
	if err != nil {
		e.log.Error(err, "[entity.GetEmojiByCode()] - cannot get emoji")
		return err
	}

	if emojis != nil {
		mochiAvatar = emojis.EmojiUrl
	}

	msg := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Description: content,
			Color:       0x62A1FE,
			Image: &discordgo.MessageEmbedImage{
				URL: image,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    fmt.Sprintf("v%s", version),
				IconURL: mochiAvatar,
			},
			Timestamp: time.Now().UTC().Format("2006-01-02 15:04:05"),
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Publish",
						Emoji: discordgo.ComponentEmoji{
							Name: "approve",
							ID:   "1077631110047080478",
						},
						Style:    discordgo.SuccessButton,
						CustomID: fmt.Sprintf(`mochi_changelog_confirm_%s`, filename),
					},
				},
			},
		},
	}

	_, err = e.discord.ChannelMessageSendComplex(e.cfg.MochiChangelogChannelID, msg)
	if err != nil {
		e.log.Error(err, "[entity.SendChangelogToChannel()] - cannot send confirm button to changelog channel")
		return err
	}

	return nil
}

func (e *Entity) ParseChangelogContent(version string, title string, content string) (string, []string) {
	replaceContent := regexp.MustCompile(`\*\*(.*?)\*\*`).ReplaceAllString(content, `\<b>$1\</b>`)
	replaceContent = strings.ReplaceAll(replaceContent, "[//]: new_line", `\newline`)
	input := []byte(strings.Split(replaceContent, "[//]: break")[0])
	reader := text.NewReader(input)
	markdownAST := goldmark.DefaultParser().Parse(reader)

	ctx := model.MarkdownContext{
		Images:         make([]string, 0),
		FirstParagraph: false,
		FirstHeading:   false,
	}

	var text string
	text = e.parseMarkDown(markdownAST, &ctx, input)
	filteredImages := make([]string, 0)
	for _, img := range ctx.Images {
		if strings.Contains(img, "imgur.com") {
			filteredImages = append(filteredImages, img)
		}
	}

	// parse strong text
	text = fmt.Sprintf("<a:gem:1095990259877158964> **%s**\n%s", title, text)
	text = regexp.MustCompile(`\\<b>(.*?)\\</b>`).ReplaceAllString(text, "**$1**")
	text = strings.ReplaceAll(text, `\newline`, "")
	text += fmt.Sprintf("\n%s [Mochi Web](%s/%s)", consts.NewChangelogDiscordFooter, consts.ChangelogUrl, version)

	return text, filteredImages
}

func (e *Entity) parseMarkDown(content ast.Node, ctx *model.MarkdownContext, source []byte) string {
	text := ""
	ast.Walk(content, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch n.Kind() {
			case ast.KindHeading:
				astHeading := n.(*ast.Heading)
				if astHeading.Level != 3 {
					text += "\n"
					text += fmt.Sprintf(`**%s**`, string(n.FirstChild().Text(source)))
				}
				ctx.FirstHeading = true
				break
			case ast.KindParagraph:
				text += "\n" + string(n.FirstChild().Text(source))
				ctx.FirstParagraph = true
				break
			case ast.KindImage:
				astImage := n.(*ast.Image)
				//text += `\<br\>`
				ctx.Images = append(ctx.Images, string(astImage.Destination))
				break
			case ast.KindListItem:
				text += "\n" + string(n.FirstChild().Text(source))
			case ast.KindLink:
				astLink := n.(*ast.Link)
				text += fmt.Sprintf(`[%s](%s)`, string(n.Text(source)), astLink.Destination)
			}
		}
		return ast.WalkContinue, nil
	})

	return text
}
