package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetEnvVariables(names ...string) (map[string]string, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
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
