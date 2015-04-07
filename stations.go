package tchou

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Fetchable interface {
	Fetch() error
}

type Station struct {
	Name, Title string

	slug string
}

func (s Station) URL() string {
	return fmt.Sprintf("http://www.sncf.com/fr/%s", s.slug)
}

func (s *Station) Fetch() error {
	_, err := goquery.NewDocument(s.URL())

	if err != nil {
		return err
	}

	// TODO
	return nil
}

type StationsList struct {
	selection   *goquery.Selection
	idx, length int
}

func (sl *StationsList) Len() int {
	return sl.length
}

func (sl *StationsList) Fetch() error {
	doc, err := goquery.NewDocument(listURL)

	if err != nil {
		return err
	}

	sl.selection = doc.Find(listStations)
	sl.length = sl.selection.Length()

	return nil
}

func (sl *StationsList) Next() *Station {
	s := sl.Get(sl.idx)
	sl.idx++
	return s
}

func (sl *StationsList) Get(idx int) *Station {

	if idx >= sl.length || idx < 0 {
		return nil
	}

	node := sl.selection.Get(idx)

	s := Station{}

	for _, attr := range node.Attr {
		switch attr.Key {
		case "href":
			s.slug = attr.Val
		case "title":
			s.Title = attr.Val
		}
	}

	return &s
}

func (sl *StationsList) GetStations() []Station {
	stations := make([]Station, sl.Len())

	for i := 0; i < sl.length; i++ {
		stations[i] = *sl.Get(i)
	}

	return stations
}
