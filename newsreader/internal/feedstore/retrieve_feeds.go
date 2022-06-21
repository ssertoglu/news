package feedstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const feedServiceEndpointUrl = "http://localhost:8090/feed"

// FeedFilter is the model which represents filtering options to search in list of feeds.
type FeedFilter struct {
	Providers  []string `json:"providers"`
	Categories []string `json:"categories"`
}

// Feed is the model which represents the info required for a feed.
type Feed struct {
	Provider string `json:"provider"`
	Category string `json:"category"`
	Url      string `json:"url"`
}

// RetrieveFeeds is the feed store client function to retrieve details of the feeds requested from the feed store service.
func (s *Client) RetrieveFeeds(filter *FeedFilter) (*[]Feed, error) {
	if filter == nil {
		return nil, errors.New("nil_reference_for_filter_argument_to_retrieve_feeds")
	}

	body, err := json.Marshal(filter)
	if err != nil {
		return nil, fmt.Errorf("error_serialising_feed_filter_to_post_to_feedstore_service: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, feedServiceEndpointUrl, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error_creating_request_for_feedstore_service_call: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error_at_feedstore_service_call: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected_status_from_feedstore_service_call: %w", err)
	}

	var feeds []Feed
	err = json.NewDecoder(resp.Body).Decode(&feeds)
	if err != nil {
		return nil, fmt.Errorf("error_unmarshalling_feedstore_response_body: %w", err)
	}

	return &feeds, nil
}
