package transport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"test.com/feedstore/internal/store"
)

type handler struct {
	service Service
}

// Service is the interface needed for feed store transport handler to call for feed functions.
type Service interface {
	SearchFeeds(filter *store.FeedFilter) (*[]store.Feed, error)
	PersistFeeds([]store.Feed) error
}

// NewHandler initialises a new handler and returns a new router instance.
func NewHandler(service Service) (http.Handler, error) {
	if service == nil {
		return nil, errors.New("nil_reference_for_service_to_initialise_handler")
	}

	handler := handler{
		service: service,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/feed", handler.searchFeeds).Methods(http.MethodPost)
	router.HandleFunc("/feed", handler.setFeeds).Methods(http.MethodPut)

	return router, nil
}
