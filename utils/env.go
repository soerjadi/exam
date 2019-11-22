package utils

import (
	"os"
	"strconv"
)

func GetEnv(env, fallback string) string {

	if value, exists := os.LookupEnv(env); exists {
		return value
	}

	return fallback
}

func GetEnvInt(env string, fallback int) int {

	if value, exists := os.LookupEnv(env); exists {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
