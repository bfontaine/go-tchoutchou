package tchou

type fetcher struct{}

func (f fetcher) Fetch(url string) {
	// TODO
}

func (f fetcher) FetchStationsList() {
	// TODO
	f.Fetch(listURL)
}
