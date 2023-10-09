package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	Error = "error >>> "
	Info  = "info >>> "
	Log   = "log >>> "
)

type Config struct {
	ClinicPath string
	BranchPath string
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("not found env")
	}

	var cfg Config

	cfg.ClinicPath = cast.ToString(getValueOrDefault("CLINIC_PATH", "./data/clinic.json"))
	cfg.BranchPath = cast.ToString(getValueOrDefault("BRANCH_PATH", "./data/branch.json"))

	return cfg
}

func getValueOrDefault(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
