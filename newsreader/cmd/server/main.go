package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"

	"test.com/newsreader/internal/articlestore"
	"test.com/newsreader/internal/feedstore"
	"test.com/newsreader/internal/service"
	"test.com/newsreader/internal/transport"
)

func main() {
	fmt.Println("starting news reader api on port 8080")

	cache := cache.New(1*time.Minute, 1*time.Minute)

	articlestore, err := articlestore.NewClient(cache)
	if err != nil {
		log.Fatal(err)
	}

	feedstore := feedstore.NewClient()

	newsService, err := service.New(feedstore, articlestore)
	if err != nil {
		log.Fatal(err)
	}

	handler, err := transport.NewHandler(newsService)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", handler))
}
