package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/honeybadger-io/honeybadger-go"
	"github.com/jinzhu/gorm"
	"github.com/viki-org/go-utils/healthcheck"
	"github.com/viki-org/go-utils/httphandlers"
	"github.com/viki-org/go-utils/queue"
	"github.com/viki-org/go-utils/queue/rabbitmq"
	"github.com/viki-org/go-utils/redis"
	"github.com/viki-org/gokit-template/pkg/svchttphandlers"

	healthRegister "github.com/viki-org/go-utils/healthcheck/register"
	sampleEndpoints "github.com/viki-org/gokit-template/pkg/endpoints/sample"
	sampleServices "github.com/viki-org/gokit-template/pkg/services/sample"
	sampleStores "github.com/viki-org/gokit-template/pkg/stores/sample"
	sampleHTTPTransport "github.com/viki-org/gokit-template/pkg/transports/http/sample"
)

// App encapsulates app dependencies.
type App struct {
	Router       *mux.Router
	DB           *gorm.DB
	Logger       log.Logger
	QueueClient  queue.MessageQueueClient
	LoggingLevel string
	Redis        redis.Redis
}

// AppConfig is the application optional config.
type AppConfig func(*App) error

// WithDB creates a db instance with its driver and connection string.
func WithDB(config Config) AppConfig {
	return func(a *App) error {
		connectionStr := fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Database.Host,
			config.Database.Port,
			config.Database.User,
			config.Database.Name,
			config.Database.Password,
		)
		db, err := gorm.Open("postgres", connectionStr)
		if err != nil {
			return err
		}
		db.DB().SetMaxIdleConns(config.Database.MaxIdleConnections)
		db.DB().SetMaxOpenConns(config.Database.MaxOpenConnections)
		db.LogMode(config.Database.LogMode)

		a.DB = db
		return nil
	}
}

// WithQueue creates a queue client to send message.
func WithQueue(config Config) AppConfig {
	return func(a *App) error {
		queueClient, err := rabbitmq.NewQueueClient(
			config.Queue.Name,
			queue.Config{
				Host:     config.Queue.Host,
				Port:     config.Queue.Port,
				Username: config.Queue.Username,
				Password: config.Queue.Password,
			},
		)
		if err != nil {
			return err
		}

		a.QueueClient = queueClient
		return nil
	}
}

// WithLoggingLevel sets the logging level for the application.
func WithLoggingLevel(config Config) AppConfig {
	return func(a *App) error {
		a.LoggingLevel = config.Server.LoggingLevel
		return nil
	}
}

// WithRedis creates a redis client for the application.
func WithRedis(config Config) AppConfig {
	return func(a *App) error {
		redisConfig := &redis.Config{
			Host:               config.Redis.Host,
			Port:               config.Redis.Port,
			Database:           config.Redis.DB,
			MaxIdleConnections: config.Redis.MaxIdleConnections,
			IdleTimeout:        config.Redis.IdleTimeout,
		}
		a.Redis = redis.NewRedisClient(redisConfig)
		return nil
	}
}

func (a *App) getLoggingLevel() level.Option {
	switch a.LoggingLevel {
	case "info":
		return level.AllowInfo()
	case "warn":
		return level.AllowWarn()
	case "error":
		return level.AllowError()
	case "debug":
		return level.AllowDebug()
	case "none":
		return level.AllowNone()
	}

	return level.AllowAll()
}

// Initialize starts
func (a *App) Initialize(configs ...AppConfig) error {
	for _, config := range configs {
		err := config(a)
		if err != nil {
			return err
		}
	}

	logger := log.NewJSONLogger(os.Stdout)
	logger = level.NewFilter(logger, a.getLoggingLevel())
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	a.Router = mux.NewRouter()
	a.Logger = logger

	a.Router.Methods("GET").Path("/ping.json").Handler(http.HandlerFunc(healthcheck.Simple))
	initializeHealthCheck(
		a.Router,
		healthRegister.NewComponentChecker("redis", healthRegister.RedisKind, a.Redis, true),
		healthRegister.NewComponentChecker("gorm", healthRegister.CustomCheckerKind, newGormChecker(a.DB), true),
		healthRegister.NewComponentChecker("rabbitMQProducer", healthRegister.CustomCheckerKind, newVikiQueueChecker(a.QueueClient), true),
	)

	options := svchttphandlers.DefaultServerOptions(a.Logger)

	sampleStore := sampleStores.NewDBStore(a.DB, a.Logger)
	sampleService := sampleServices.NewAPIService(sampleStore, a.Logger)
	sampleEndpoint := sampleEndpoints.MakeServerEndpoints(sampleService)

	a.Router.Methods("GET").Path("/sample_route.json").Handler(httptransport.NewServer(
		sampleEndpoint.RetrieveEndpoint,
		sampleHTTPTransport.DecodeRetrieveRequest,
		svchttphandlers.EncodeStatusOKResponse,
		options...,
	))

	return nil
}

func initializeHealthCheck(router *mux.Router, components ...healthRegister.CheckComponent) {
	handler := healthRegister.SetupComponents(components...)
	router.Methods("GET").Path("/v4/healthcheck.json").Handler(handler)
}

// Run starts the application at a particular port.
func (a *App) Run(config Config) {
	a.Logger.Log("transport", "HTTP", "port", config.Server.Port)

	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "staging" {
		honeybadger.Configure(honeybadger.Configuration{
			APIKey: config.Server.HoneybadgerAPIKey,
			Env:    os.Getenv("ENV"),
		})
		defer honeybadger.Monitor()
		httphandlers.StartServer(strconv.Itoa(config.Server.Port), honeybadger.Handler(a.Router))
	} else {
		httphandlers.StartServer(strconv.Itoa(config.Server.Port), a.Router)
	}
}
