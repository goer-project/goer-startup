package config

import (
	"goer-startup/internal/pkg/log"
	"goer-startup/pkg/db"
)

var (
	Cfg *Config
)

type Config struct {
	Server Server           `mapstructure:"server" json:"server" yaml:"server"`
	Mysql  *db.MySQLOptions `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Log    *log.Options     `mapstructure:"log" json:"log" yaml:"log"`
}

type Server struct {
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}
