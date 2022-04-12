package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// FileReader read config from file
type FileReader struct {
	filename string
	dirname  string
}

// NewFileLoader create new file loader with filename and dirname
func NewFileLoader(filename, dirname string) Loader {
	return &FileReader{filename, dirname}
}

// Load from yml file
func (r *FileReader) Load(v viper.Viper) (*viper.Viper, error) {
	filePath := fmt.Sprintf("%s/%s", r.dirname, r.filename)
	err := godotenv.Load(filePath)
	if err != nil {
		return nil, err
	}
	v.AutomaticEnv()
	return &v, nil
}
