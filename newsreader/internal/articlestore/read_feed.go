package articlestore

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ungerik/go-rss"

	"test.com/newsreader/internal/feedstore"
)

const dateFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

// readFeed is the function reads from the feed url and unmarshalls response to an article array.
func readFeed(feed *feedstore.Feed) ([]Article, error) {
	if feed == nil {
		return nil, errors.New("nil_reference_for_feed_argument_to_read_feed")
	}

	feedResponse, err := http.Get(feed.Url)
	if err != nil {
		return nil, fmt.Errorf("error_reading_feed: %w", err)
	}

	channel, err := rss.Regular(feedResponse)
	if err != nil {
		return nil, fmt.Errorf("error_unmarshalling_feed_response: %w", err)
	}

	log.Printf("Feed %v read with %v articles.\n", feed.Url, len(channel.Item))

	articles := make([]Article, len(channel.Item))

	for i, item := range channel.Item {
		time, err := item.PubDate.Parse()

		if err != nil {
			time, err = item.PubDate.ParseWithFormat(dateFormat)
			if err != nil {
				return nil, fmt.Errorf("error_parsing_date_published_of_feed_item '%s': %w", item.GUID, err)
			}
		}

		var thumbNail string

		if len(item.Enclosure) > 0 {
			thumbNail = item.Enclosure[0].URL
		}

		articles[i] = Article{
			Provider:      feed.Provider,
			Category:      feed.Category,
			Title:         item.Title,
			Description:   item.Description,
			Link:          item.Link,
			ThumbNailURL:  thumbNail,
			DatePublished: time,
		}
	}

	return articles, nil
}
