package search

import (
	"fmt"
)

type Repository struct {
	InMemorySearchService SearchService
}

func NewRepository(inMemorySearchService InMemorySearchService) Repository {
	return Repository{
		InMemorySearchService: inMemorySearchService,
	}
}

func (r Repository) PerformSearch(query string) ([]string, error) {
	result := r.InMemorySearchService.FindMatches(query)

	if len(result) == 0 {
		return []string{}, fmt.Errorf("no match found for %s", query)
	}

	return result, nil
}

type SearchService interface {
	FindMatches(query string) []string
}
