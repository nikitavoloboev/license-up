package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app         = kingpin.New("license-up", "A command-line tool to make licences.")
	mit         = app.Command("mit", "Create MIT license.")
	mitName     = mit.Arg("name", "Name of license holder.").Required().String()
	mitSurname  = mit.Arg("surname", "Surname of license holder.").Required().String()
	mitWebsite  = mit.Arg("website", "Website of license holder").String()
	bsd2        = app.Command("bsd2", "Create BSD 2-Clause license.")
	bsd2Name    = bsd2.Arg("name", "Name of license holder.").Required().String()
	bsd2Surname = bsd2.Arg("surname", "Surname of license holder.").Required().String()
	bsd3        = app.Command("bsd", "Create BSD 3-Clause license.")
	bsd3Name    = bsd3.Arg("name", "Name of license holder.").Required().String()
	bsd3Surname = bsd3.Arg("surname", "Surname of license holder.").Required().String()
	cc0         = app.Command("cc0", "Create CC0 license.")
	unlicense   = app.Command("unlicense", "Create Unlicense license.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case mit.FullCommand():
		mitCreateWithSite(string(*mitName), string(*mitSurname), string(*mitWebsite))
	case bsd2.FullCommand():
		bsd2Create(string(*bsd2Name), string(*bsd2Surname))
	case bsd3.FullCommand():
		bsd3Create(string(*bsd3Name), string(*bsd3Surname))
	case cc0.FullCommand():
		cc0Create()
	case unlicense.FullCommand():
		unlicenseCreate()
	}
}
