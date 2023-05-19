package search

type Searcher struct {
}

var result = "Item 1"

func (s Searcher) Search(query string) {
	result = "Item 1"
}

func (s Searcher) GetResult() string {
	return result
}
