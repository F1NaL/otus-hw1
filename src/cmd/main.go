package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"os/signal"
	"otus-hw-1/src/httpd"
	"syscall"
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
	waitForKillSignal(logger)
}

func waitForKillSignal(logger *zap.Logger) {
	sysKillSignal := make(chan os.Signal, 1)
	signal.Notify(sysKillSignal, os.Interrupt, syscall.SIGTERM)
	s := <-sysKillSignal
	logger.Info("got system signal", zap.String("signal: ", s.String()))
}
