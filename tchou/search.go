package tchou

import (
	"errors"

	"github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
)

var (
	ErrNotImplemented = errors.New("Not implemented")
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

// FindStation tries to find a station by its name or its code.
func FindStation(name string) (*Station, error) {
	// TODO
	return nil, ErrNotImplemented
}
