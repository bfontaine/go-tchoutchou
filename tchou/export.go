package tchou

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

var csvHeaders = []string{"name", "address", "lat", "long", "ref"}

func (s *Station) records() []string {
	return []string{
		s.Name,
		s.Address,
		strconv.FormatFloat(s.Lat, 'g', -1, 64),
		strconv.FormatFloat(s.Long, 'g', -1, 64),
		s.Ref,
	}
}

func (l *StationsList) WriteCSV(writer io.Writer) error {
	w := csv.NewWriter(writer)

	if err := w.Write(csvHeaders); err != nil {
		return err
	}

	err := l.each(func(s Station) error {
		return w.Write(s.records())
	})

	if err != nil {
		return err
	}

	w.Flush()
	return w.Error()
}

func (l *StationsList) WriteJSON(writer io.Writer) error {
	w := json.NewEncoder(writer)

	return w.Encode(l.GetStations())
}

func (l *StationsList) Print(writer io.Writer) error {
	return l.each(func(s Station) error {
		fmt.Printf("%s\n", s.Name)
		return nil
	})
}
