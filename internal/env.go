package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvVariables(names ...string) (map[string]string, error) {
	m, err := loadEnvVariables(names...)
	if err != nil {
		log.Print("was not able to load all required env variables: ", err)
		log.Print("loading .env file and retrying")
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
		m, err = loadEnvVariables(names...)
	}
	return m, nil
}

func loadEnvVariables(names ...string) (map[string]string, error) {
	m := make(map[string]string)
	for _, name := range names {
		val, ok := os.LookupEnv(name)
		if !ok {
			return nil, fmt.Errorf("environment variable $%s is not set", name)
		}
		m[name] = val
	}
	return m, nil
}
