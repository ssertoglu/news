package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/go-playground/validator"
)

// Persist saves the given feed data a file in JSON format
func (s *Store) Persist(feeds []Feed) error {
	if len(feeds) == 0 {
		return errors.New("empty_feed_list_is_not_valid")
	}

	validate := validator.New()

	for _, feed := range feeds {
		if err := validate.Struct(feed); err != nil {
			return fmt.Errorf("error_validating_feed: %w", err)
		}

		_, err := url.ParseRequestURI(feed.Url)
		if err != nil {
			return fmt.Errorf("invalid_url_fo_feed: '%s'", feed.Url)
		}
	}

	file, err := json.MarshalIndent(feeds, "", " ")
	if err != nil {
		return fmt.Errorf("error_marshalling_feed_data: %w", err)
	}

	err = ioutil.WriteFile(feedLocation, file, 0644)
	if err != nil {
		return fmt.Errorf("error_saving_feed_data: %w", err)
	}

	s.feeds = feeds

	return nil
}
