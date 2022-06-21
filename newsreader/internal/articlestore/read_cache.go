package articlestore

import (
	"errors"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"

	"test.com/newsreader/internal/feedstore"
)

//Cache is the interface needed for readcache function to read from and write to a cache.
type Cache interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
}

// readCache checks if the artciles for a feed is in the cache.
// If exists return directly from cache.
// If not calls feed reader and caches and returns response from the read.
func (c *Client) readCache(feed *feedstore.Feed) (*[]Article, error) {
	if feed == nil {
		return nil, errors.New("nil_reference_for_feed_argument_to_read_cache")
	}

	if cachedValue, found := c.cache.Get(feed.Url); found {
		return cachedValue.(*[]Article), nil
	} else {
		articles, err := readFeed(feed)
		if err != nil {
			return nil, fmt.Errorf("error_reading_feed_from_url '%s': %w", feed.Url, err)
		}

		c.cache.Set(feed.Url, &articles, cache.DefaultExpiration)

		return &articles, nil
	}
}
