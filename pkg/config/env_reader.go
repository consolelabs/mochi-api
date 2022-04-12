package config

import (
	"github.com/spf13/viper"
)

// ENVReader load config from env
type ENVReader struct {
}

// NewENVLoader create new env loader
func NewENVLoader() Loader {
	return &ENVReader{}
}

// Load env into viper
func (r *ENVReader) Load(v viper.Viper) (*viper.Viper, error) {
	v.AutomaticEnv()
	return &v, nil
}
