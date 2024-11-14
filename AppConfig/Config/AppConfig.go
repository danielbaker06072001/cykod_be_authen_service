package Config

import (
	"log"
	"os"

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
	}

	return &config, nil
}

func Connect(cfg *Appconfig) (*gorm.DB, error) {
	dsn := cfg.Postgres.DB_URL

	db_data, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database data")
		return nil, err
	}

	log.Println("Connected to database data")
	return db_data, nil
}
