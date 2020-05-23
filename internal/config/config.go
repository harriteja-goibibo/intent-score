package config

import (
	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

// New returns new viper config by reading the file in filepath.
func New(filepath string) (config *viper.Viper, err error) {

	config = viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(filepath)
	if err = config.ReadInConfig(); err != nil {
		err = errors.Wrap(err, "NewConfig")
		return
	}

	config.WatchConfig()
	return
}
