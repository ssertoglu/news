package articlestore

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"test.com/newsreader/internal/feedstore"
)

// Article is the model which represents article present in the feed.
type Article struct {
	Provider      string    `json:"provider"`
	Category      string    `json:"category"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Link          string    `json:"link"`
	ThumbNailURL  string    `json:"thumbNailURL"`
	DatePublished time.Time `json:"datePublished"`
}

// RetrieveArticles is the service handler to return articles fro the feed requested.
func (c *Client) RetrieveArticles(feeds *[]feedstore.Feed) (*[]Article, error) {
	if feeds == nil {
		return nil, errors.New("nil_reference_for_feeds_argument_to_retrieve_articles")
	}

	articles := []Article{}

	for _, feed := range *feeds {
		feedArticles, err := c.readCache(&feed)
		if err != nil {
			return nil, fmt.Errorf("error_checking_cache_for_feed_at_url '%s': %w", feed.Url, err)
		}

		articles = append(articles, *feedArticles...)
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].DatePublished.Unix() > articles[j].DatePublished.Unix()
	})

	return &articles, nil
}
