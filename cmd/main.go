package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // for development
	"github.com/rs/cors"
	"github.com/sidaurukdedi/go-boiler/config"
	"github.com/sidaurukdedi/go-boiler/internal/user"
	"github.com/sidaurukdedi/go-boiler/pkg/mongodb"
	"github.com/sidaurukdedi/go-boiler/pkg/pubsub"
	"github.com/sidaurukdedi/go-boiler/pkg/response"
	"github.com/sidaurukdedi/go-boiler/pkg/server"
	"github.com/sidaurukdedi/go-boiler/pkg/validator"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
)

var (
	tracer       *apm.Tracer
	cfg          *config.Config
	indexMessage string = "Application is running properly"
)

func init() {
	tracer = apm.DefaultTracer
	cfg = config.Load()
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(cfg.Logger.Formatter)
	logger.SetReportCaller(true)
	// logger.AddHook(&ddlogrus.DDContextLogHook{})
	// logger.AddHook(hook.NewStdoutLoggerHook(logrus.New(), cfg.Logger.Formatter))

	// set crypto
	// aesGcm := crypto.NewAES256GCM(cfg.Crypto.Secret, cfg.Crypto.Pepper)

	// set validator
	vld := validator.NewValidator()

	// set mongodb
	mca := mongodb.NewClientAdapter(cfg.Mongodb.ClientOptions)
	if err := mca.Connect(context.Background()); err != nil {
		logger.Fatal(err)
	}
	mdb := mca.Database(cfg.Mongodb.Database)

	// set mariadb source object
	// dbRW, err := sql.Open(cfg.MariadbRW.Driver, cfg.MariadbRW.DSN)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// if err := dbRW.Ping(); err != nil {
	// 	logger.Fatal(err)
	// }
	// dbRW.SetConnMaxLifetime(time.Minute * 3)
	// dbRW.SetMaxOpenConns(cfg.MariadbRW.MaxOpenConnections)
	// dbRW.SetMaxIdleConns(cfg.MariadbRW.MaxIdleConnections)

	// set mariadb read only object
	// dbRO, err := sql.Open(cfg.MariadbRO.Driver, cfg.MariadbRO.DSN)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// if err := dbRO.Ping(); err != nil {
	// 	logger.Fatal(err)
	// }
	// dbRO.SetConnMaxLifetime(time.Minute * 3)
	// dbRO.SetMaxOpenConns(cfg.MariadbRO.MaxOpenConnections)
	// dbRO.SetMaxIdleConns(cfg.MariadbRO.MaxIdleConnections)

	// set redis object
	// rdb := rv8.NewClient(cfg.Redis.Options)
	// if _, err := rdb.Ping(context.Background()).Result(); err != nil {
	// 	logger.Fatal(err)
	// }

	// set publisher object
	saramaAsyncProducer, err := sarama.NewAsyncProducer(
		cfg.SaramaKafka.Addresses,
		cfg.SaramaKafka.Config,
	)

	if err != nil {
		logger.Fatal(err)
	}

	publisher := pubsub.NewSaramaKafkaProducerAdapter(logger, &pubsub.SaramaKafkaProducerAdapterConfig{
		AsyncProducer: saramaAsyncProducer,
	})

	router := mux.NewRouter()
	router.HandleFunc("/go-boiler", index)

	userRepository := user.NewUserRepository(logger, mdb, "user-tester")
	userUsecase := user.NewUserUsecase(user.UserUsecaseProperty{
		ServiceName:    cfg.Application.Name,
		Logger:         logger,
		Location:       cfg.Application.Location,
		UserRepository: userRepository,
	})

	user.NewUserHTTPHandler(logger, vld, router, userUsecase)

	// set cors
	handler := cors.New(cors.Options{
		AllowedOrigins:   cfg.Application.AllowedOrigins,
		AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	// initiate server
	srv := server.NewServer(logger, handler, cfg.Application.Port)
	srv.Start()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm

	srv.Close()
	publisher.Close()
	// dbRW.Close()
	// dbRO.Close()
	mca.Disconnect(context.Background())

}

func index(w http.ResponseWriter, r *http.Request) {
	resp := response.NewSuccessResponse(nil, response.StatOK, indexMessage)
	response.JSON(w, resp)
}
