package service

import "errors"

type NewsService struct {
	feedRetriever    FeedRetriever
	articleRetriever ArticleRetriever
}

// New returns a new service instance.
func New(feedRetriever FeedRetriever, articleRetriever ArticleRetriever) (*NewsService, error) {
	if feedRetriever == nil {
		return nil, errors.New("nil_reference_for_feedRetriever_to_initialise_service_instance")
	}

	if articleRetriever == nil {
		return nil, errors.New("nil_reference_for_articleRetriever_to_initialise_service_instance")
	}

	newsService := NewsService{
		feedRetriever:    feedRetriever,
		articleRetriever: articleRetriever,
	}

	return &newsService, nil
}
