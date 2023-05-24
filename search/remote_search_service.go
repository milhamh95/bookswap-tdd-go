package search

type RemoteSearchService struct {
	SearchApi SearchApi
}

func NewRemoteSearchService(searchApi SearchApi) RemoteSearchService {
	return RemoteSearchService{
		SearchApi: searchApi,
	}
}

func (r RemoteSearchService) FindMatches(query string) ([]string, error) {
	return r.SearchApi.FindMatches(query)
}
