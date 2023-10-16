package environment

import (
	"fmt"
	"os"
)

func GetWithFallback(key string, fallback string) string {
	if value, success := os.LookupEnv(key); success {
		return value
	}
	return fallback
}

func GetEnvOrDie(key string) string {
	value := os.Getenv(key)

	if value == "" {
		err := fmt.Errorf("Missing environment variable %s", key)
		panic(err)
	}

	return value
}
