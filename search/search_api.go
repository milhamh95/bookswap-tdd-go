package search

type SearchApi interface {
	FindMatches(query string) ([]string, error)
}
