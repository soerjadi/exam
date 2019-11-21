package utils

import (
	"os"
)

func GetEnv(env, fallback string) string {

	if value, exists := os.LookupEnv(env); exists {
		return value
	}

	return fallback
}
