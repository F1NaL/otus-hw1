package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"otus-hw-1/src/httpd"
)

func main() {
	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := logConfig.Build()
	if err != nil {
		log.Fatalln(err.Error())
	}

	logger.Info("service started")
	//ctx := grace.ShutdownContext(context.Background())

	go httpd.HttpServer(logger)
}
