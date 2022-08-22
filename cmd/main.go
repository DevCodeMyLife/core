package main

import (
	"context"
	"core/internal/repositories"
	"core/internal/services"
	"core/internal/transport/http/controller"
	"fmt"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/driver/postgres"

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

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow`,
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PORT"),
	)

	logger.Debug(dsn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Notice("App Run ...")

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	repo := repositories.NewRepositories(db)
	serv := services.NewServices(repo)

	go controller.NewController(ctx, serv)

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
