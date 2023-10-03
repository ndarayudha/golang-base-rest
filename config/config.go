package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server ServerConfig
	Mysql  MysqlConfig
	Logger Logger
}

// ServerConfig config struct
type ServerConfig struct {
	AppVersion   string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Mode         string
}

// MysqlConfig config struct
type MysqlConfig struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDBName   string
	MysqlDriver   string
	MysqlNetwork  string
}

// Logger config struct
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
