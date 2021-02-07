package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.pie.apple.com/someorg/somemodule/config"
)

func main() {
	cfg := &config.Config{}
	envconfig.MustProcess("", cfg)
}
