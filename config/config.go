package config

import (
	"fmt"

	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/util"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type config struct {
	Postgres postgres `yaml:"POSTGRES"`
	APP      app      `yaml:"APP"`
}

//CFG is config instance
var CFG config

//LoadConfig loads and initializes config list
func LoadConfig(configPath string) *models.Configuration {

	viper.SetEnvPrefix("EDUCATIVE")
	viper.AddConfigPath(".")
	viper.SetConfigFile(configPath)
	err := viper.MergeInConfig()
	if err != nil {
		fmt.Println("Error in reading config")
		panic(err)
	}

	err = viper.Unmarshal(&CFG, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
	})
	if err != nil {
		fmt.Println("Error in un-marshaling config")
		panic(err)
	}
	// fillBanningJobTime()
	if CFG.APP.LogLevel == "info" {
		fmt.Printf("%#v \n", CFG)
	}

	postgres, err := GetPostgres()
	if err != nil {
		panic(err)
	}
	util.Initialize()
	return &models.Configuration{
		PostgresConnection: postgres,
	}
}
