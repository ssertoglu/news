package service

import (
	"errors"

	"test.com/feedstore/internal/store"
)

// FeedService is the struct to represent a feed service functions.
type FeedService struct {
	store Store
}

// Store is the interface the service will need to interact with data layer.
type Store interface {
	Search(filter *store.FeedFilter) (*[]store.Feed, error)
	Persist(feeds []store.Feed) error
}

// New returns a new feed service instance.
func New(store Store) (*FeedService, error) {
	if store == nil {
		return nil, errors.New("nil_reference_for_store_to_initialise_service_instance")
	}

	return &FeedService{
		store: store,
	}, nil
}
