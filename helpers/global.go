package helpers

import (
	"os"
)

func Getenv(key, fallback string) string {
	if value, err := os.LookupEnv(key); err {
		return value
	}
	return fallback
}