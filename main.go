package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app          = kingpin.New("license-up", "A command-line tool to make licences.").Version("1.0.0")
	force        = app.Flag("force", "Create a license even if license already exists.").Short('f').Bool()
	md           = app.Flag("md", "Create a LICENSE.md file instead of a LICENSE file.").Bool()
	mit          = app.Command("mit", "Create MIT license.")
	mitName      = mit.Arg("name", "Name of license holder.").Required().String()
	mitSurname   = mit.Arg("surname", "Surname of license holder.").Required().String()
	mitWebsite   = mit.Arg("website", "Website of license holder").String()
	bsd2         = app.Command("bsd2", "Create BSD 2-Clause license.")
	bsd2Name     = bsd2.Arg("name", "Name of license holder.").Required().String()
	bsd2Surname  = bsd2.Arg("surname", "Surname of license holder.").Required().String()
	bsd3         = app.Command("bsd3", "Create BSD 3-Clause license.")
	bsd3Name     = bsd3.Arg("name", "Name of license holder.").Required().String()
	bsd3Surname  = bsd3.Arg("surname", "Surname of license holder.").Required().String()
	cc0          = app.Command("cc0", "Create CC0 license.")
	unlicense    = app.Command("unlicense", "Create Unlicense license.")
	gpl2         = app.Command("gpl2", "Create GNU General Public License version 2.")
	gpl3         = app.Command("gpl3", "Create GNU General Public License version 3.")
	isc          = app.Command("isc", "Create ISC license.")
	iscName      = isc.Arg("name", "Name of license holder.").Required().String()
	iscSurname   = isc.Arg("surname", "Surname of license holder.").Required().String()
	wtfpl        = app.Command("wtfpl", "Create WTFPL license.")
	wtfplName    = wtfpl.Arg("name", "Name of license holder.").Required().String()
	wtfplSurname = wtfpl.Arg("surname", "Surname of license holder.").Required().String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	// Check to see if the user wants to create a `LICENSE.md` file instead of a `LICENSE` file
	fileName := "LICENSE"
	if bool(*md) == true {
		fileName = "LICENSE.md"
	}
	// Check to see if we are overwriting any existing license files
	if bool(*force) == false {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(os.Stdin)
		overwrite := false
		hasLicense := false
		for _, f := range files {
			if f.Name() == fileName {
				hasLicense = true
				fmt.Print("There is already a license present in current directory. Do you want to overwrite it with a new one? [y/N] ")
				text, _ := reader.ReadString('\n')
				switch text {
				case "y\n":
					overwrite = true
				case "Y\n":
					overwrite = true
				}
			}
		}
		if hasLicense == true && overwrite == false {
			os.Exit(0)
		}
	}
	// Create the licenses
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case mit.FullCommand():
		if string(*mitWebsite) == "" {
			mitCreate(string(*mitName), string(*mitSurname), fileName)
		} else {
			mitCreateWithSite(string(*mitName), string(*mitSurname), string(*mitWebsite), fileName)
		}
	case bsd2.FullCommand():
		bsd2Create(string(*bsd2Name), string(*bsd2Surname), fileName)
	case bsd3.FullCommand():
		bsd3Create(string(*bsd3Name), string(*bsd3Surname), fileName)
	case cc0.FullCommand():
		cc0Create(fileName)
	case unlicense.FullCommand():
		unlicenseCreate(fileName)
	case gpl2.FullCommand():
		gpl2Create(fileName)
	case gpl3.FullCommand():
		gpl3Create(fileName)
	case isc.FullCommand():
		iscCreate(string(*iscName), string(*iscSurname), fileName)
	case wtfpl.FullCommand():
		wtfplCreate(string(*wtfplName), string(*wtfplSurname), fileName)
	}
}
