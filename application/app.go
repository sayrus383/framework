package application

import (
	"github.com/joho/godotenv"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/logger"
)

type app struct {
	name string
	env  string

	cfg Config
	log logger.Logger
}

func New() App {
	var (
		cfg = newFlagConfig()
		log logger.Logger
	)

	err := godotenv.Load(".env")
	if err != nil {
		defer func() {
			log.Warn("unable to load .env file, loading system config")
		}()
	}

	var appName, env, logLevel string

	cfg.StringVar(&appName, "FRAMEWORK_APPNAME", "new-app", "app name")
	cfg.StringVar(&env, "FRAMEWORK_APPENV", "development", "app environment")
	cfg.StringVar(&logLevel, "FRAMEWORK_LOGLEVEL", "DEBUG", "app log level")

	log = logger.NewSlogLogger(appName, env, logger.StringToLevel(logLevel))

	return &app{
		name: appName,
		env:  env,
		cfg:  cfg,
		log:  log,
	}
}

func (a app) Env() string {
	return a.env
}

func (a app) Name() string {
	return a.name
}

func (a app) Run(start func(a App) error) error {
	a.log.Info("app starting...")
	defer a.log.Info("app started")

	return start(a)
}

func (a app) GetConfig() Config {
	return a.cfg
}

func (a app) GetLogger() logger.Logger {
	return a.log
}
