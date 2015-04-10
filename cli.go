package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/alecthomas/kingpin"
	"github.com/bfontaine/go-tchoutchou/tchou"
)

var (
	app = kingpin.New("tchou", "A command-line interface to French railroad stations")

	// subcommands
	stationCmd = app.Command("station", "show infos about a station")
	listCmd    = app.Command("list", "list all stations")

	// args
	stationRef = stationCmd.Arg("ref", "station ref").Required().String()

	// flags
	csvList  = listCmd.Flag("csv", "print the list as CSV").Bool()
	jsonList = listCmd.Flag("json", "print the list as JSON").Bool()
)

func main() {
	app.Version("0.1.0")

	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch parsed {
	case stationCmd.FullCommand():
		s, err := tchou.FindStationByRef(*stationRef)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s (%s)\n\n%s\n(%f, %f)\n",
			s.Name, s.Ref, s.Address, s.Lat, s.Long)

	case listCmd.FullCommand():
		l, err := tchou.GlobalList()

		if err != nil {
			log.Fatal(err)
		}

		out := os.Stdout

		if *csvList {
			l.WriteCSV(out)
		} else if *jsonList {
			l.WriteJSON(out)
		} else {
			l.Print(out)
		}
	}
}
