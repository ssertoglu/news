package store

import "errors"

// Search filters the feed data with given parameters.
// If no parameter specified complete list of feeds will be returned.
func (s *Store) Search(filter *FeedFilter) (*[]Feed, error) {
	if filter == nil {
		return nil, errors.New("nil_reference_for_filter_argument_to_search_feeds")
	}

	if len(filter.Categories) > 0 || len(filter.Providers) > 0 {
		selectedFeeds := []Feed{}
		for _, feed := range s.feeds {
			var selected bool

			for _, category := range filter.Categories {
				if category == feed.Category {
					selected = true
					break
				}
			}

			if !selected {
				for _, provider := range filter.Providers {
					if provider == feed.Provider {
						selected = true
						break
					}
				}
			}

			if selected {
				selectedFeeds = append(selectedFeeds, feed)
			}
		}
		return &selectedFeeds, nil
	} else {
		return &s.feeds, nil
	}
}
