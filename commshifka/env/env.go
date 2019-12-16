package env

import (
	"os"
)

func GetEnv(key string, defaultKey string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultKey
	}
	return val
}
