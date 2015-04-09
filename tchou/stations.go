package tchou

import (
	"fmt"
	"strings"

	"github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
)

const (
	listStations = ".container-stations a.sncfcom-colors-internal-automatic"
)

// A Station is a train station
type Station struct {
	Name, Address string

	Lat, Long float64

	// Ref is a short unique reference to a station
	Ref string

	slug string
}

// URL is the URL of the station
func (s Station) URL() string {
	return fmt.Sprintf("http://www.sncf.com/fr/%s", s.slug)
}

// fetch fetches the station's page and populate its fields
func (s *Station) fetch() error {
	doc, err := goquery.NewDocument(s.URL())

	if err != nil {
		return err
	}

	desc := doc.Find("#description")

	s.Name = desc.Find("label").Text()
	s.Address = desc.Find("input[name=adresseGare]").Text()
	s.Lat = softParseFloat(desc.Find("input[name=adresseLat]").Text())
	s.Long = softParseFloat(desc.Find("input[name=adresseLong]").Text())

	parts := strings.Split(s.slug, "-")
	lparts := len(parts)
	if lparts > 1 {
		s.Ref = parts[lparts-1]
	}

	// note: we could also parse the next trains

	return nil
}

// A StationsList is a list of stations
type StationsList struct {
	selection   *goquery.Selection
	idx, length int
}

// Len returns the number of stations in the list
func (sl *StationsList) Len() int {
	return sl.length
}

func (sl *StationsList) populateFromDoc(doc *goquery.Document) {
	sl.selection = doc.Find(listStations)
	sl.length = sl.selection.Length()
}

// Next returns one station at a time. It is more effective when you need to
// process each station independently.
func (sl *StationsList) Next() Station {
	s := sl.Get(sl.idx)
	sl.idx++
	return s
}

// More checks if there are more stations available with Next()
func (sl *StationsList) More() bool {
	return sl.idx < sl.length-1
}

// Get returns the station at the given index in the list
func (sl *StationsList) Get(idx int) (s Station) {
	node := sl.selection.Get(idx)

	for _, attr := range node.Attr {
		switch attr.Key {
		case "href":
			s.slug = attr.Val
		case "title":
			s.Name = attr.Val
		}
	}

	return s
}

// GetStations return a slice with all stations at once
func (sl *StationsList) GetStations() []Station {
	stations := make([]Station, sl.Len())

	for i := 0; i < sl.length; i++ {
		stations[i] = sl.Get(i)
	}

	return stations
}
