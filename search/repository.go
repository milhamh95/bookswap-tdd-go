package search

import "fmt"

type Repository struct {
	AvailableQueries map[string]string
}

func NewRepository() Repository {
	return Repository{
		AvailableQueries: map[string]string{
			"item":    "Item 1",
			"another": "Another Item",
		},
	}
}

func (r Repository) PerformSearch(query string) string {
	item, ok := r.AvailableQueries[query]
	if !ok {
		return fmt.Sprintf("No match found for %s", query)
	}

	return item
}
