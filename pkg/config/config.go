package config

import "time"

var Cfg config

type config struct {
	App      App      `yaml:"app"`
	XUI      XUI      `yaml:"xui"`
	Telegram Telegram `yaml:"telegram"`
}

type App struct {
}

type XUI struct {
	Address  string        `yaml:"address"`
	Username string        `yaml:"username"`
	Password string        `yaml:"password"`
	Timeout  time.Duration `yaml:"timeout"`
}

type Telegram struct {
	Token string `yaml:"token"`
}
