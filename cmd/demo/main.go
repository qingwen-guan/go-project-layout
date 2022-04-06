package main

import (
	"flag"

	"github.com/qingwen-guan/go-project-layout/internal/config"
	service "github.com/qingwen-guan/go-project-layout/internal/service/subservice"
	"github.com/qingwen-guan/go-project-layout/internal/telemetry"
)

func main() {
	confFilePath := flag.String("config", "", "toml config file")
	flag.Parse()

	conf := config.NewConfigFromFile(*confFilePath)
	logger := telemetry.NewZapLogger(conf)

	logger.Info("init")

	demo := service.NewDemo(conf, logger)
	demo.Run()
}
