package config

import (
	"log"
	"os"
)

func GetEnvOrFatal(kev string) string {
	val := os.Getenv(kev)
	if val == "" {
		log.Fatalf("env value \"%s\" not set", kev)
	}

	return val
}

func GetEnvOrDefault(key, def string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}

	return def
}
