package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DemoDSN string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}
