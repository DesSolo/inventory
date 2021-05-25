package main

import (
	"fmt"
	"inventory/internal/config"
	"inventory/internal/server"
	"inventory/internal/server/storage"
	"log"
)

func getStorage(storageName string) (storage.Storage, error) {
	switch storageName {
	case "memory":
		return storage.NewMemoryStorage(), nil
	default:
		return nil, fmt.Errorf("storage \"%s\" not supported", storageName)
	}
}

func main() {
	listenAddress := config.GetEnvOrFatal("LISTEN_ADDRESS")
	token := config.GetEnvOrFatal("TOKEN")
	persistantStorage, err := getStorage(
		config.GetEnvOrDefault("STORAGE_TYPE", "memory"),
	)
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewRestServer(token, persistantStorage)
	log.Printf("server listen at %s", listenAddress)
	if err := server.Run(listenAddress); err != nil {
		log.Fatal(err)
	}
}
