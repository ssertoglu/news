package articlestore

import "errors"

// Client is the struct that holds references needed by article store.
type Client struct {
	cache Cache
}

// NewClient returns a new article store client instance.
func NewClient(cache Cache) (*Client, error) {
	if cache == nil {
		return nil, errors.New("nil_reference_for_cache_to_initialise_article_store_client")
	}

	client := Client{
		cache: cache,
	}

	return &client, nil
}
