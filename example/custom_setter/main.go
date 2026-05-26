package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Port      string `env:"PORT"`
	JWTSecret string `env:"JWT_SECRET"`
	Roles     roles  `env:"ROLES"`
}

type roles []string

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

func (r *roles) SetValue(s string) error { _ = "STUB: not implemented"; return nil }

func setEnvValues() error { _ = "STUB: not implemented"; return nil }
