package config

import (
	"github.com/spf13/viper"
)

var Conf *Env

type Env struct {
	Database Database `mapstructure:"database"`
	Server   Server   `mapstructure:"server"`
	Jwt      Jwt      `mapstructure"jwt"`
}
type Database struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
type Server struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Jwt struct {
	Secret  string `mapstructure:"secret"`
	Expiry  uint   `mapStructure:"expiry"`
	Refresh uint   `mapStrcuture:"refresh"`
}

func GetEnv() *Env {
	if Conf == nil {
		v := viper.New()
		v.SetConfigType("yaml")
		v.SetConfigFile("./.yaml")
		err := v.ReadInConfig()

		if err != nil {
			panic(err)
		}

		err = v.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	}

	return Conf
}
