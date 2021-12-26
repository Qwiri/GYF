package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.SetHandler(cli.Default)
}

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	sc := make(chan os.Signal, 1)

	go func(cancel chan os.Signal) {
		if err := app.Listen(":8080"); err != nil {
			log.WithError(err).Error("Starting WebServer failed")
		}
		sc <- os.Interrupt
	}(sc)

	signal.Notify(sc, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sc

	log.Info("Shutting down WebServer")
	if err := app.Shutdown(); err != nil {
		log.WithError(err).Error("Shutting down failed")
	}
	log.Info("Done!")
}
