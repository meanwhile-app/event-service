package configs

import (
	"log"

	"github.com/joho/godotenv"
)

var appEnv map[string]string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		println("1111")
		log.Fatal(err)
	}

	env, err := godotenv.Read()
	if err != nil {
		println("2222")
		log.Fatal(err)
	}
	appEnv = env
}

func GetEnv() map[string]string {
	return appEnv
}
