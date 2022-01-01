package main

import (
	"github.com/Qwiri/GYF/backend/internal/server"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	recov "github.com/gofiber/fiber/v2/middleware/recover"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	JanitorTime      = 30 * time.Second
	JanitorCleanTime = 24 * time.Hour
)

var (
	LimiterWhitelist = []string{
		"/game",
		"/game/",
	}
	DevMode bool
)

func init() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	if !strings.HasPrefix(os.Getenv("BUILD"), "prod") {
		DevMode = true
	}
}

func main() {
	svr := server.NewServer(DevMode)
	_ = svr.RouteCreateGame(nil) // TODO: remove dummy game

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	var corsConfig = cors.ConfigDefault
	if DevMode {
		log.Info("Enabling monitor")
		app.Use(monitor.New())
	} else {
		// limit requests to only 10 per Minute
		app.Use(limiter.New(limiter.Config{
			Max:        10,
			Expiration: time.Minute,
			Next: func(c *fiber.Ctx) bool {
				for _, a := range LimiterWhitelist {
					if a == c.Path() {
						return true
					}
				}
				return false
			},
		}))

		// allow specific origins only
		corsConfig = cors.Config{
			AllowOrigins: "https://prod.gyf.d2a.io, https://staging.gyf.d2a.io",
		}
	}

	app.Use(cors.New(corsConfig))

	// log requests
	app.Use(logger.New())

	app.Use(recov.New())

	// create API routes
	svr.CreateRoutes(app)

	// start janitor
	go func(s *server.GYFServer) {
		for {
			time.Sleep(JanitorTime)

			log.Info("[Janitor] Checking ...")
			svr.JanitorCheck(JanitorCleanTime)
			log.Info("[Janitor] Done!")
		}
	}(svr)

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
