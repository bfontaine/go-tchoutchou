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

	// flags
	csvList  = listCmd.Flag("csv", "print the list as CSV").Bool()
	jsonList = listCmd.Flag("json", "print the list as JSON").Bool()
)

func main() {
	app.Version("0.1.0")

	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch parsed {
	case stationCmd.FullCommand():
		// TODO
	case listCmd.FullCommand():
		// TODO

		// just a test for now
		l, err := tchou.GlobalList()

		if err != nil {
			log.Fatal(err)
		}

		for l.More() {
			s := l.Next()
			fmt.Printf("%s\n", s.Name)
		}
	}

	// TODO
}
