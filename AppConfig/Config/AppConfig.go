package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Appconfig struct {
	Server struct {
		Host    string `env:"HOST"`
		GinPort string `env:"GIN_PORT"`
	}

	Postgres struct {
		DB_URL string `env:"DB_URL"`
	}

	Redis struct { 
		Addr     string `env:"REDIS_ADDR"`
		Password string `env:"REDIS_PASSWORD"`
		DB   	 int 	`env:"REDIS_DB"`
	}
}

var env string

func SetEnvironment(environment string) {
	env = environment
}

func LoadConfig() (*Appconfig, error) {
	if env != ".env" && env != "local.env" {
		env = env + ".env"
	}

	err := godotenv.Load(env)
	if err != nil {
		log.Println("Error loading .env file")
		return nil, err
	}

	config := Appconfig{
		Server: struct {
			Host    string `env:"HOST"`
			GinPort string `env:"GIN_PORT"`
		}{
			Host:    os.Getenv("HOST"),
			GinPort: os.Getenv("GIN_PORT"),
		},

		Postgres: struct {
			DB_URL string `env:"DB_URL"`
		}{
			DB_URL: os.Getenv("DB_URL"),
		},

		Redis : struct { 
			Addr     string `env:"REDIS_ADDR"`
			Password string `env:"REDIS_PASSWORD"`
			DB   	 int 	`env:"REDIS_DB"`
		} { 
			Addr: os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB: 0,
		},
	}

	return &config, nil
}

func Connect(cfg *Appconfig) (*gorm.DB, error) {
	dsn := cfg.Postgres.DB_URL

	db_data, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database data")
	return db_data, nil
}

func ConnectRedis(cfg *Appconfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to redis")
	}
	log.Println("Connected to redis")
	return client, nil
}