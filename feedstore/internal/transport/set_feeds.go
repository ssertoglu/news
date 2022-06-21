package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"test.com/feedstore/internal/store"
)

// setFeeds is the transport layer handler to update persisted feeds.
func (h *handler) setFeeds(w http.ResponseWriter, r *http.Request) {
	var feeds []store.Feed
	err := json.NewDecoder(r.Body).Decode(&feeds)
	if err != nil {
		fmt.Println(err)
	}

	err = h.service.PersistFeeds(feeds)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
}
