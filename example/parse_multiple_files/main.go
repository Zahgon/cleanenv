package main

import (
	"log"
)

type config struct {
	Port         string         `yaml:"port"`
	JWTSecret    string         `yaml:"jwt_secret"`
	Database     databaseConfig `yaml:"database"`
	EmailService emailService   `yaml:"email_service"`
}

type databaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"ssl_mode"`
}

type emailService struct {
	EmailKey                    string `yaml:"email_key"`
	AccountValidationTemplateID string `yaml:"account_validation_template_id"`
}

func main() {
	cfg, err := ParseConfigFiles("./db_config.yaml", "./email_config.yaml", "./general_config.yaml")
	if err != nil {
		log.Printf("Error parsing config files: %v", err)
		return
	}

	log.Println("Parsed Configuration")
	log.Println(*cfg)

	return
}

func ParseConfigFiles(files ...string) (*config, error) { _ = "STUB: not implemented"; return nil, nil }
