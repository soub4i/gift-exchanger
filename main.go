package main

import (
	"os"

	"github.com/soub4i/giftsxchanger/pkg/api"
	"github.com/soub4i/giftsxchanger/pkg/datastore"
)

func main() {
	version := os.Getenv("API_VERION")

	if version == "" {
		version = "v1"
	}

	ds := datastore.GetDS()
	ds.Seed()
	api := api.Register(version)

	api.Run("0.0.0.0:8080")

}
