package service

import "test.com/feedstore/internal/store"

// SearchFeeds is the service handler for search request.
// Should normally do more like abstraction of data model but things simplified here.
func (s *FeedService) SearchFeeds(filter *store.FeedFilter) (*[]store.Feed, error) {
	return s.store.Search(filter)
}
