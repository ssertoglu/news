package transport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"test.com/newsreader/internal/articlestore"
	"test.com/newsreader/internal/feedstore"
)

type handler struct {
	articleLoader ArticleLoader
}

// ArticleLoader is the interface needed for article loader handler to use to load news from the service.
type ArticleLoader interface {
	LoadArticles(filter *feedstore.FeedFilter) (*[]articlestore.Article, error)
}

// NewHandler initialises a new handler and returns a new router instance.
func NewHandler(articleLoader ArticleLoader) (http.Handler, error) {
	if articleLoader == nil {
		return nil, errors.New("nil_reference_for_artice_loader_to_initialise_service_handler")
	}

	handler := handler{
		articleLoader: articleLoader,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/news", handler.loadArticles).Methods(http.MethodPost)

	return router, nil
}
