package utils

import (
	"fmt"
	"os"
)

func GetEnvOrErr(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("missing %s environment variable", key)
	}
	return value, nil
}
