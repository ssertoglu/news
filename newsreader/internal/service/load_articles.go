package service

import (
	"errors"
	"fmt"

	"test.com/newsreader/internal/articlestore"
	"test.com/newsreader/internal/feedstore"
)

type FeedRetriever interface {
	RetrieveFeeds(filter *feedstore.FeedFilter) (*[]feedstore.Feed, error)
}

type ArticleRetriever interface {
	RetrieveArticles(feeds *[]feedstore.Feed) (*[]articlestore.Article, error)
}

// LoadArticles is the service handler to load articles requested.
func (s *NewsService) LoadArticles(filter *feedstore.FeedFilter) (*[]articlestore.Article, error) {
	if filter == nil {
		return nil, errors.New("nil_reference_for_filter_argument_to_load_articles")
	}

	feeds, err := s.feedRetriever.RetrieveFeeds(filter)
	if err != nil {
		return nil, fmt.Errorf("error_retrieving_feeds_at_feed_retriever: %w", err)
	}

	articles, err := s.articleRetriever.RetrieveArticles(feeds)
	if err != nil {
		return nil, fmt.Errorf("error_retrieving_articled_at_article_retriever: %w", err)
	}

	return articles, nil
}
