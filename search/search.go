package search

type Searcher struct {
}

var result = "Item 1"

func (s Searcher) Search(query string) {
	if query == "item" {
		result = "Item 1"
	} else if query == "another" {
		result = "Another Item"
	}

}

func (s Searcher) GetResult() string {
	return result
}
