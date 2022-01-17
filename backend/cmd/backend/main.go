package main

import (
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Qwiri/GYF/backend/internal/handlers"
	"github.com/Qwiri/GYF/backend/internal/server"
	"github.com/Qwiri/GYF/backend/pkg/config"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recov "github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	JanitorTime      = 30 * time.Second
	JanitorCleanTime = 24 * time.Hour
	ConfigFile       = "config.toml"
)

var (
	LimiterWhitelist = []string{
		"/game",
		"/game/",
	}
	DevMode bool
)

func init() {
	log.SetHandler(text.Default)
	log.SetLevel(log.DebugLevel)

	if !strings.HasPrefix(os.Getenv("BUILD"), "prod") {
		DevMode = true
	}
}

func readConfig() {
	// check if config already exists
	_, err := os.Stat(ConfigFile)
	if err == nil {
		// parse config
		if _, err = toml.DecodeFile(ConfigFile, config.Obj); err != nil {
			log.WithError(err).Warn("cannot read/decode config")
			return
		}
	}
	// (re-) write config
	var file *os.File
	if file, err = os.Create(ConfigFile); err != nil {
		log.WithError(err).Warnf("cannot create config")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.WithError(err).Warn("cannot close config")
		}
	}(file)
	if err = toml.NewEncoder(file).Encode(config.Obj); err != nil {
		log.WithError(err).Warn("cannot encode config")
	}
}

func initConfig() {
	val := reflect.ValueOf(config.Obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Type().Implements(reflect.TypeOf((*config.Initer)(nil)).Elem()) {
			_ = field.MethodByName("Init").Call(nil)
		}
	}
}

func main() {
	log.Infof("Backend-Version: %s.%s:%s", config.Version, config.GitBranch, config.GitCommit)

	readConfig()
	initConfig()

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	var corsConfig = cors.ConfigDefault
	if !DevMode {
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

		// "proper" random seed
		rand.Seed(time.Now().Unix())
	}
	app.Use(cors.New(corsConfig))

	// log requests
	app.Use(logger.New())

	app.Use(recov.New())

	// create gyf server and create routes
	svr := server.NewServer(DevMode)
	svr.CreateRoutes(app)

	// start janitor
	go func(s *server.GYFServer) {
		for {
			time.Sleep(JanitorTime)

			svr.JanitorCheck(JanitorCleanTime)
			handlers.JanitorChatChats()
		}
	}(svr)

	sc := make(chan os.Signal, 1)
	go func(cancel chan os.Signal) {
		if err := app.Listen(":8080"); err != nil {
			log.WithError(err).Error("Starting WebServer failed")
		}
		sc <- os.Interrupt
	}(sc)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	log.Info("Shutting down WebServer")
	if err := app.Shutdown(); err != nil {
		log.WithError(err).Error("Shutting down failed")
	}
	log.Info("Done!")
}
