package config

import (
	"github.com/spf13/viper"
	"log"
	"github.com/joho/godotenv"
  "os"
  "fmt"
)

type Config struct {
	Server struct {
		Port string
	}

	App struct {
	    Name string
      Env string
      LogLevel string
      DB DBConfig
	}

// following MSSQL is not required its replaced with .env
	MSSQL struct {
		ConnString string
	}

	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func Load() (*Config, error) {

    // Load .env file
    err := godotenv.Load()
    if err != nil {
      log.Println("No .env file found, relying on environment variables")
    }

    // Load config.yml
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".") // root
    if err := viper.ReadInConfig(); err != nil {
      log.Fatalf("Error reading config.yml: %v", err)
      return nil, err
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
      return nil, err
    }

    log.Println("Loaded configuration")

    cfg.App.DB = DBConfig{
    	User:     os.Getenv("DB_USER"),
    	Password: os.Getenv("DB_PASSWORD"),
    	Host:     os.Getenv("DB_HOST"),
    	Port:     os.Getenv("DB_PORT"),
    	Name:     os.Getenv("DB_NAME"),
    }
    return &cfg, nil
}

// Build full connection string
func (db DBConfig) ConnString() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		db.User, db.Password, db.Host, db.Port, db.Name)
}
