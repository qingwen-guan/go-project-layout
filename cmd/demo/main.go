package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/qingwen-guan/go-project-layout/internal/config"
	service "github.com/qingwen-guan/go-project-layout/internal/service/subservice"
	"github.com/qingwen-guan/go-project-layout/internal/telemetry"
)

func main() {
	confFilePath := flag.String("config", "", "toml config file")
	flag.Parse()

	conf, err := config.NewConfigFromFile(*confFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode config file %s, err=%s\n", *confFilePath, err.Error())
		os.Exit(1)
	}
	logger := telemetry.NewZapLogger(conf)
	defer logger.Sync()

	logger.Info("init")

	demo := service.NewDemo(conf, logger)
	demo.Run()
}
