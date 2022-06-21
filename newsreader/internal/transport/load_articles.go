package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"test.com/newsreader/internal/feedstore"
)

// loadArticles is the transport layer handler for news article requests.
func (h *handler) loadArticles(w http.ResponseWriter, r *http.Request) {
	var filter feedstore.FeedFilter

	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		fmt.Println(err)
	}

	articles, err := h.articleLoader.LoadArticles(&filter)
	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(articles)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
