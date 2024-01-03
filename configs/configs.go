package configs

import (
	"log"

	"github.com/joho/godotenv"
)

var appEnv map[string]string

func LoadEnv() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	env, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	appEnv = env
	return env
}

func GetEnv() map[string]string {
	return appEnv
}
