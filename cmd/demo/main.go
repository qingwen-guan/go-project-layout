package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/qingwen-guan/go-project-layout/internal/config"
	service "github.com/qingwen-guan/go-project-layout/internal/service/subservice"
	"github.com/qingwen-guan/go-project-layout/internal/telemetry"
	"go.uber.org/zap"
)

func calcClientName(logger *zap.Logger) string {
	u, err := user.Current()
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	return "DEBUG-" + u.Username
}

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

	clientName := calcClientName(logger)

	logger.Info("init", zap.String("client_name", clientName))

	demo := service.NewDemo(conf, logger)
	demo.Run()
}
