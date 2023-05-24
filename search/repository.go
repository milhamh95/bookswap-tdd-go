package search

import (
	"fmt"
)

type Repository struct {
	searchService SearchService
}

func NewRepository(searchService SearchService) Repository {
	return Repository{
		searchService: searchService,
	}
}

func (r Repository) PerformSearch(query string) ([]string, error) {
	result, err := r.searchService.FindMatches(query)
	if err != nil {
		return []string{}, err
	}

	if len(result) == 0 {
		return []string{}, fmt.Errorf("no match found for %s", query)
	}

	return result, nil
}

type SearchService interface {
	FindMatches(query string) ([]string, error)
}
