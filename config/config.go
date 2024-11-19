package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

// Config is an app configuration.
type Config struct {
	Application struct {
		Port           string
		Name           string
		AllowedOrigins []string
		Location       *time.Location
	}
	BasicAuth struct {
		Username string
		Password string
	}
	Crypto struct {
		Secret string
		Pepper string
	}
	Logger struct {
		Formatter logrus.Formatter
	}
	SaramaKafka struct {
		Addresses []string
		Config    *sarama.Config
	}
	Mongodb struct {
		ClientOptions *options.ClientOptions
		Database      string
	}
}

// Load will load the configuration.
func Load() *Config {
	cfg := new(Config)
	cfg.app()
	cfg.basicAuth()
	cfg.crypto()
	cfg.logFormatter()
	cfg.sarama()
	cfg.mongodb()
	return cfg
}

func (cfg *Config) app() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	appName := os.Getenv("APP_NAME")
	port := os.Getenv("APP_PORT")
	timezone := os.Getenv("APP_TIMEZONE")
	if l, err := time.LoadLocation(timezone); err == nil {
		loc = l
	}
	rawAllowedOrigins := strings.Trim(os.Getenv("APP_ALLOWED_ORIGINS"), " ")

	allowedOrigins := make([]string, 0)
	if rawAllowedOrigins == "" {
		allowedOrigins = append(allowedOrigins, "*")
	} else {
		allowedOrigins = strings.Split(rawAllowedOrigins, ",")
	}

	cfg.Application.Port = port
	cfg.Application.Name = appName
	cfg.Application.AllowedOrigins = allowedOrigins
	cfg.Application.Location = loc
}

func (cfg *Config) basicAuth() {
	username := os.Getenv("BASIC_AUTH_USERNAME")
	password := os.Getenv("BASIC_AUTH_PASSWORD")

	cfg.BasicAuth.Username = username
	cfg.BasicAuth.Password = password
}

func (cfg *Config) crypto() {
	secret := os.Getenv("AES_SECRET")
	pepper := os.Getenv("AES_PEPPER")

	cfg.Crypto.Pepper = pepper
	cfg.Crypto.Secret = secret
}

func (cfg *Config) logFormatter() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			// _, filename := path.Split(f.File)
			filename := fmt.Sprintf("%s:%d", f.File, f.Line)
			return funcname, filename
		},
	}

	cfg.Logger.Formatter = formatter
}

func (cfg *Config) sarama() {
	brokers := os.Getenv("KAFKA_BROKERS")
	sslEnable, _ := strconv.ParseBool(os.Getenv("KAFKA_SSL_ENABLE"))
	username := os.Getenv("KAFKA_USERNAME")
	password := os.Getenv("KAFKA_PASSWORD")
	ca := strings.Replace(os.Getenv("KAFKA_CA_ROOT"), `\n`, "\n", -1)
	clientCert := strings.Replace(os.Getenv("KAFKA_CLIENT_CERT"), `\n`, "\n", -1)
	clientKey := strings.Replace(os.Getenv("KAFKA_CLIENT_KEY"), `\n`, "\n", -1)

	sc := sarama.NewConfig()
	sc.Version = sarama.V2_1_0_0
	if username != "" {
		sc.Net.SASL.User = username
		sc.Net.SASL.Password = password
		sc.Net.SASL.Enable = true
	}
	sc.Net.TLS.Enable = sslEnable

	if clientCert != "" {
		tlscfg := tls.Config{}
		cert, _ := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
		tlscfg.Certificates = []tls.Certificate{cert}

		pool := x509.NewCertPool()
		if ok := pool.AppendCertsFromPEM([]byte(ca)); !ok {
			fmt.Println("kafka fail append pem")
		}

		tlscfg.RootCAs = pool
		sc.Net.TLS.Config = &tlscfg
	}

	// consumer config
	sc.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	sc.Consumer.Offsets.Initial = sarama.OffsetOldest

	// producer config
	sc.Producer.Retry.Backoff = time.Millisecond * 500

	cfg.SaramaKafka.Addresses = strings.Split(brokers, ",")
	cfg.SaramaKafka.Config = sc
}

func (cfg *Config) mongodb() {
	appName := os.Getenv("APP_NAME")
	uri := os.Getenv("MONGODB_URL")
	db := os.Getenv("MONGODB_DATABASE")
	minPoolSize, _ := strconv.ParseUint(os.Getenv("MONGODB_MIN_POOL_SIZE"), 10, 64)
	maxPoolSize, _ := strconv.ParseUint(os.Getenv("MONGODB_MAX_POOL_SIZE"), 10, 64)
	maxConnIdleTime, _ := strconv.ParseInt(os.Getenv("MONGODB_MAX_IDLE_CONNECTION_TIME_MS"), 10, 64)

	opts := options.Client().
		ApplyURI(uri).
		SetAppName(appName).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize).
		SetMaxConnIdleTime(time.Millisecond * time.Duration(maxConnIdleTime)).
		SetMonitor(apmmongo.CommandMonitor())

	cfg.Mongodb.ClientOptions = opts
	cfg.Mongodb.Database = db

}
