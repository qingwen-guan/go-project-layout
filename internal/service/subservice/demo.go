package service

import (
	"github.com/qingwen-guan/go-project-layout/internal/config"
	"go.uber.org/zap"
)

type Demo struct {
	conf   *config.Config
	logger *zap.Logger
}

func NewDemo(conf *config.Config, logger *zap.Logger) *Demo {
	return &Demo{
		conf:   conf,
		logger: logger,
	}
}

func (demo *Demo) Run() {
	logger := demo.logger

	logger.Info("running demo")
}
