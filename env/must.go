package env

import (
	"errors"
	"os"
)

var (
	errNotFound    = errors.New("the environment variable is not found")
	errEmptyString = errors.New("the environment variable is an empty string")
)

func MustNonEmpty(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(errNotFound)
	}
	if len(val) == 0 {
		panic(errEmptyString)
	}

	return val
}

func Or(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
