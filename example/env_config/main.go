package main

import (
	"log"
	"net/url"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Port      string        `env:"PORT"`
	JWTSecret string        `env:"JWT_SECRET"`
	DB        url.URL       `env:"DB"`
	Start     time.Time     `env:"START"`
	TTL       time.Duration `env:"TTL" env-required:"true"`
}

func main() {
	err := setEnvValues()
	if err != nil {
		panic(err)
	}

	var cfg config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	log.Println("Parsed Configuration")
	log.Println(cfg)
	return
}

func setEnvValues() error { _ = "STUB: not implemented"; return nil }
