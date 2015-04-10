package tchou

import (
	"errors"
	"fmt"

	"github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
)

// TODO check http://www.sncf.com/theme/js/cityList.js

// GlobalList returns the list of all stations
func GlobalList() (*StationsList, error) {
	sl := StationsList{}

	doc, err := goquery.NewDocument(listURL)

	if err != nil {
		return nil, err
	}

	sl.populateFromDoc(doc)

	return &sl, nil
}

// FindStation returns a station given its ref
func FindStationByRef(ref string) (*Station, error) {

	s := Station{slug: fmt.Sprintf("gare-d-%s", ref), Ref: ref}

	err := s.fetch()

	return &s, err
}
