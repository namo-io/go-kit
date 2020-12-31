package config

import (
	"context"
	"path"
	"strings"

	"github.com/namo-io/go-kit/log"

	"github.com/spf13/viper"
)

// New Config Object
func New(ctx context.Context, filepath string) (*viper.Viper, error) {
	logger := log.New().WithContext(ctx)

	dir, file := path.Split(filepath)
	if dir == "" {
		dir = "./"
	}
	ext := path.Ext(filepath)
	filename := strings.TrimSuffix(file, ext)

	if len(ext) < 2 {
		logger.Fatalf("check config file path: %s", filepath)
	}

	conf := viper.New()
	conf.AddConfigPath(dir)
	conf.SetConfigType(ext[1:])
	conf.SetConfigName(filename)
	conf.SetConfigFile(filepath)

	// Find and read the config file
	if err := conf.ReadInConfig(); err != nil {
		logger.Fatalf("Error reading config file, %s", err)
		return nil, err
	}
	// Confirm which config file is used
	logger.Infof("Config: %s\n", conf.ConfigFileUsed())

	return conf, nil
}
