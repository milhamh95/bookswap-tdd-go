package search

import "fmt"

type Repository struct {
	AvailableQueries map[string]string
}

func NewRepository(availableQueries map[string]string) Repository {
	return Repository{
		AvailableQueries: availableQueries,
	}
}

func (r Repository) PerformSearch(query string) string {
	item, ok := r.AvailableQueries[query]
	if !ok {
		return fmt.Sprintf("No match found for %s", query)
	}

	return item
}
