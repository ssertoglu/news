package main

import (
	"fmt"
	"log"
	"net/http"

	"test.com/feedstore/internal/service"
	"test.com/feedstore/internal/store"
	"test.com/feedstore/internal/transport"
)

func main() {
	fmt.Println("starting feedstore api on port 8090")

	store, err := store.New()
	if err != nil {
		log.Fatal(err)
	}

	feedService, err := service.New(store)
	if err != nil {
		log.Fatal(err)
	}

	handler, err := transport.NewHandler(feedService)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8090", handler))
}
