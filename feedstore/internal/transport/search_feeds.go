package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"test.com/feedstore/internal/store"
)

func (h *handler) searchFeeds(w http.ResponseWriter, r *http.Request) {
	var filter store.FeedFilter
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		fmt.Println(err)
	}

	feeds, err := h.service.SearchFeeds(&filter)
	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(feeds)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
