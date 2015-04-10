package tchou

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
)

const (
	listStations = ".container-stations a.sncfcom-colors-internal-automatic"
)

var (
	ErrStationNotFound = errors.New("Station not found")
)

var httpClient = http.DefaultClient

// A Station is a train station
type Station struct {
	Name    string `json:"name"`
	Address string `json:"address"`

	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`

	// Ref is a short unique reference to a station
	Ref string `json:"ref"`

	slug string
}

// URL is the URL of the station
func (s Station) URL() string {
	return fmt.Sprintf("http://www.sncf.com/fr/%s", s.slug)
}

func (s Station) request() *http.Request {
	r, _ := http.NewRequest("GET", s.URL(), nil)
	r.AddCookie(&http.Cookie{Name: "has_js", Value: "1"})
	return r
}

// fetch fetches the station's page and populate its fields
func (s *Station) fetch() error {
	resp, err := httpClient.Do(s.request())

	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		return ErrStationNotFound
	}

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		return err
	}

	desc := doc.Find("#description")

	s.Name = desc.Find("label").Text()
	s.Address, _ = desc.Find("input[name=adresseGare]").Attr("value")

	lat, _ := desc.Find("input[name=adresseLat]").Attr("value")
	long, _ := desc.Find("input[name=adresseLong]").Attr("value")

	s.Lat = softParseFloat(lat)
	s.Long = softParseFloat(long)

	// note: we could also parse the next trains

	return nil
}

func newStation(name, slug string) Station {
	s := Station{slug: slug, Name: name}

	parts := strings.Split(s.slug, "-")
	lparts := len(parts)
	if lparts > 1 {
		s.Ref = parts[lparts-1]
	}

	return s
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
func (sl *StationsList) Get(idx int) Station {
	node := sl.selection.Get(idx)

	var slug, name string

	for _, attr := range node.Attr {
		switch attr.Key {
		case "href":
			slug = attr.Val
		case "title":
			name = attr.Val
		}
	}

	return newStation(name, slug)
}

func (sl *StationsList) each(fn func(Station) error) (err error) {
	for sl.More() {
		if err = fn(sl.Next()); err != nil {
			return
		}
	}

	return
}

// GetStations return a slice with all stations at once
func (sl *StationsList) GetStations() []Station {
	stations := make([]Station, sl.Len())

	for i := 0; i < sl.length; i++ {
		stations[i] = sl.Get(i)
	}

	return stations
}
