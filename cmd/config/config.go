package config

import (
	"encoding/json"
	"os"

	dopplergoruntime "github.com/ayatmaulana/doppler-go-runtime"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"

	"github.com/tekpriest/poprev/cmd/utils"
)

type Config struct {
	Env   string `env:"ENV"`
	Port  int    `env:"PORT"         envDefault:"3000"`
	DBURL string `env:"DATABASE_URL"`
	RedisConfig
}

type RedisConfig struct {
	RedisHost string `env:"REDIS_HOST"`
	RedisPass string `env:"REDIS_PASS"`
	RedisUser string `env:"REDIS_USER"`
	RedisPort int    `env:"REDIS_PORT"`
	RedisDB   int    `env:"REDIS_DB"   envDefault:"0"`
}

func LoadConfig(flag string) *Config {
	config := Config{}

	environment := bootstrapConfigWithProvider(flag)

	utils.PanicOnError(
		env.Parse(&config, env.Options{
			Environment:     environment,
			RequiredIfNoDef: false,
		}))

	return &config
}

func bootstrapConfigWithProvider(flag string) map[string]string {
	// if we are in prod, use doppler
	var res map[string]string

	if flag == "prod" {
		doppler := dopplergoruntime.NewDopplerRuntime(dopplergoruntime.DopplerRuntimeConfig{
			Token:   os.Getenv("DOPPLER_TOKEN"),
			Project: os.Getenv("DOPPLER_PROJECT"),
			Config:  os.Getenv("DOPPLER_CONFIG"),
		})
		utils.PanicOnError(doppler.Load())

		secrets, err := doppler.DownloadSecret()
		if err != nil {
			utils.PanicOnError(err, "there was an error downloading secrets")
		}
		utils.PanicOnError(json.Unmarshal(secrets, &res))
	}
	if flag == "dev" {
		secrets, err := godotenv.Read()
		if err != nil {
			utils.PanicOnError(err, "error reading env file. exiting")
		}
		res = secrets
	}
	return res
}
