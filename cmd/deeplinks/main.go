package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cloudberrybot/deeplinks/postgres"
	"github.com/cloudberrybot/deeplinks/web"
)

func main() {
	d := os.Getenv("DATABASE_URL")

	if d == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	log.Println("Connected to database")

	store, err := postgres.NewStore(d)
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)

	log.Println("Listening on :3000")

	http.ListenAndServe(":3000", h)
}
