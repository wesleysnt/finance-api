package config

import (
	"github.com/spf13/viper"
)

type Env struct {
	Database Database `mapstructure:"database"`
	Server   Server   `mapstructure:"server"`
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

func GetEnv() *Env {
	env := &Env{}
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile("./.yaml")
	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(env)
	if err != nil {
		panic(err)
	}

	return env
}
