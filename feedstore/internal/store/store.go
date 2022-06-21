package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const feedLocation = "../../feeds.json"

// FeedFilter is the model which represents filtering options to search in list of feeds.
type FeedFilter struct {
	Providers  []string `json:"providers"`
	Categories []string `json:"categories"`
}

// Feed is the model which represents the info required for a feed.
type Feed struct {
	Provider string `json:"provider" validate:"required"`
	Category string `json:"category" validate:"required"`
	Url      string `json:"url" validate:"required"`
}

// Store is the struct to represent a data store for feeds.
type Store struct {
	feeds []Feed
}

// load reads from the persisted JSON file where feed data is stored into an array.
// If file does not exists an empty array will be returned.
func load() ([]Feed, error) {
	if _, err := os.Stat(feedLocation); err != nil {
		if os.IsNotExist(err) {
			return []Feed{}, nil
		}
	}

	file, err := os.Open(feedLocation)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	var feeds []Feed
	json.Unmarshal(byteValue, &feeds)
	defer file.Close()

	return feeds, nil
}

// New returns a new feed store instance.
func New() (*Store, error) {
	feeds, err := load()
	if err != nil {
		fmt.Println(err)
	}

	return &Store{
		feeds: feeds,
	}, nil
}
