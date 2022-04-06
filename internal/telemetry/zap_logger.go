package telemetry

import (
	"os"

	"github.com/qingwen-guan/go-project-layout/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewZapLogger(conf *config.Config) *zap.Logger {
	zapConfig := zap.NewProductionEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

	newCore := func(filename string, level zapcore.Level) zapcore.Core {
		return zapcore.NewCore(
			zapcore.NewJSONEncoder(zapConfig),
			zapcore.AddSync(&lumberjack.Logger{
				Filename:  filename,
				MaxSize:   50, // megabytes
				LocalTime: true,
			}),
			level,
		)
	}

	var cores []zapcore.Core
	for _, logDir := range conf.LogDirs {
		switch logDir {
		case "stdout":
			core := zapcore.NewCore(
				zapcore.NewJSONEncoder(zapConfig),
				zapcore.AddSync(os.Stdout),
				zapcore.InfoLevel,
			)
			cores = append(cores, core)

		default:
			cores = append(cores, newCore(logDir+"/"+conf.AppName+".debug", zapcore.DebugLevel))
			cores = append(cores, newCore(logDir+"/"+conf.AppName+".info", zapcore.InfoLevel))
			cores = append(cores, newCore(logDir+"/"+conf.AppName+".error", zapcore.WarnLevel))
		}
	}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}
