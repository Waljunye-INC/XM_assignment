package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	envLoaded *options
}

type options struct {
	DBHost        string `env:"DB_HOST,required"`
	DBPort        string `env:"DB_PORT,required"`
	DBName        string `env:"DB_NAME,required"`
	DBUsername    string `env:"DB_USER,required"`
	DBPassword    string `env:"DB_PASSWORD,required"`
	PublicApiPort int    `env:"PUBLIC_API_PORT,required"`
}

func LoadFromEnv(fallbackFile *string) (cfg *Config, err error) {
	cfg = &Config{envLoaded: &options{}}

	if fallbackFile != nil {
		err = godotenv.Load(*fallbackFile)
	}
	*cfg.envLoaded, err = env.ParseAs[options]()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) DBHost() string {
	return cfg.envLoaded.DBHost
}

func (cfg *Config) DBPort() string {
	return cfg.envLoaded.DBPort
}

func (cfg *Config) DBName() string {
	return cfg.envLoaded.DBName
}

func (cfg *Config) DBUsername() string {
	return cfg.envLoaded.DBUsername
}

func (cfg *Config) DBPassword() string {
	return cfg.envLoaded.DBPassword
}

func (cfg *Config) PublicApiPort() int {
	return cfg.envLoaded.PublicApiPort
}
