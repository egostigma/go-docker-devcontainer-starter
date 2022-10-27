package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		DB   `yaml:"db"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// DB -.
	DB struct {
		Host      string `env-required:"true" yaml:"db_host"     env:"DB_HOST"`
		Port      string `env-required:"true" yaml:"db_port"     env:"DB_PORT"`
		User      string `env-required:"true" yaml:"db_user"     env:"DB_USER"`
		Password  string `env-required:"true" yaml:"db_password" env:"DB_PASSWORD"`
		Name      string `env-required:"true" yaml:"db_name"     env:"DB_NAME"`
		ParseTime string `yaml:"db_parse_time" env:"DB_PARSE_TIME"`
		Loc       string `yaml:"db_loc"        env:"DB_LOC"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	errEnv := godotenv.Load()
	if errEnv == nil {
		cfg.ReadFromDotEnv()
	}

	return cfg, nil
}

func (c *Config) ReadFromDotEnv() {
	if os.Getenv("APP_NAME") != "" {
		c.App.Name = os.Getenv("APP_NAME")
	}
	if os.Getenv("PORT") != "" {
		c.HTTP.Port = os.Getenv("PORT")
	}
	if os.Getenv("DB_HOST") != "" {
		c.DB.Host = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_PORT") != "" {
		c.DB.Port = os.Getenv("DB_PORT")
	}
	if os.Getenv("DB_USER") != "" {
		c.DB.User = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		c.DB.Password = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("DB_NAME") != "" {
		c.DB.Name = os.Getenv("DB_NAME")
	}
	if os.Getenv("DB_PARSE_TIME") != "" {
		c.DB.ParseTime = os.Getenv("DB_PARSE_TIME")
	}
	if os.Getenv("DB_LOC") != "" {
		c.DB.Loc = os.Getenv("DB_LOC")
	}
}
