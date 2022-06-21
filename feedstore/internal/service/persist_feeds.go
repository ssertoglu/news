package service

import "test.com/feedstore/internal/store"

// SearchFeeds is the service handler for search request.
// Should normally do more like abstraction of data model but things simplified here.
func (s *FeedService) PersistFeeds(feeds []store.Feed) error {
	return s.store.Persist(feeds)
}
