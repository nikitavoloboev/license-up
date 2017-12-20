package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("license-up", "A command-line tool to make licences.")
	mit        = app.Command("mit", "Create MIT license.")
	mitName    = mit.Arg("name", "Name of license holder.").Required().String()
	mitSurname = mit.Arg("surname", "Surname of license holder.").Required().String()
	mitWebsite = mit.Arg("website", "Website of license holder").String()

	// TODO: add license
	// cc1 = app.Command("cc1", "Create CC1 license.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case mit.FullCommand():
		mitCreateWithSite(string(*mitName), string(*mitSurname), string(*mitWebsite))
	}
}
