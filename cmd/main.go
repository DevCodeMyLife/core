package main

import (
	"context"
	"core/internal/transport/http/controller"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ivahaev/go-logger"
)

const (
	shutDownDuration = 5 * time.Second
)

func main() {
	err := logger.SetLevel("debug")
	if err != nil {
		panic(fmt.Sprintf(`failed init logs, error: %s`, err.Error()))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Notice("App Run ...")

	go controller.NewController(&ctx)

	<-GracefulShutdown()
	_, forceCancel := context.WithTimeout(ctx, shutDownDuration)

	logger.Notice("Graceful Shutdown")

	defer forceCancel()
}

func GracefulShutdown() chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	return done
}
