package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

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
	boml         = app.Command("boml", "Create Blue Oak Model license.")
)

func readInJSON(file string) map[string]interface{} {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	if err := json.Unmarshal(byteValue, &result); err != nil {
		panic(err)
	}

	return result
}

func replaceText(file string, text1 string, text2 string) error {
	read, err := ioutil.ReadFile(file)
	newContents := strings.Replace(string(read), text1, text2, -1)
	err = ioutil.WriteFile(file, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
	return nil
}

func createLicense(licenseName string, fileName string, name string, surname string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	licenseFile := dir + "/licenses/" + licenseName + "/" + fileName
	cmd := exec.Command("cp", licenseFile, dir)
	cmd.Stdout = os.Stdout
	if err = cmd.Run(); err != nil {
		if fileName == "LICENSE" {
			fmt.Println("This license hasn't been added yet.\nYou can add it via a pull request!")
		}
		if fileName == "LICENSE.md" {
			fmt.Println("The markdown version of this license hasn't been added yet.\nYou can add it via a pull request!")
		}
		log.Fatal(err)
	}
	placeholders := readInJSON(dir + "/placeholders.json")
	license := placeholders[licenseName].(map[string]interface{})
	year, _, _ := time.Now().Date()
	for key, value := range license {
		text := value.(string)
		if text != "" {
			if string(key) == "name" {
				replaceText(dir+"/"+fileName, text, name+" "+surname)
			}
			if string(key) == "year" {
				replaceText(dir+"/"+fileName, text, strconv.Itoa(year))
			}
		}
	}
	fmt.Println("License was created")
	return nil
}

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
			createLicense("mit", fileName, string(*mitName), string(*mitSurname))
		} else {
			mitCreateWithSite(string(*mitName), string(*mitSurname), string(*mitWebsite), fileName)
		}
	case bsd2.FullCommand():
		createLicense("bsd2", fileName, string(*bsd2Name), string(*bsd2Surname))
	case bsd3.FullCommand():
		createLicense("bsd3", fileName, string(*bsd3Name), string(*bsd3Surname))
	case cc0.FullCommand():
		createLicense("cc0", fileName, "", "")
	case unlicense.FullCommand():
		createLicense("unlicense", fileName, "", "")
	case gpl2.FullCommand():
		createLicense("gpl2", fileName, "", "")
	case gpl3.FullCommand():
		createLicense("gpl3", fileName, "", "")
	case isc.FullCommand():
		createLicense("isc", fileName, string(*iscName), string(*iscSurname))
	case wtfpl.FullCommand():
		createLicense("wtfpl", fileName, string(*wtfplName), string(*wtfplSurname))
	case boml.FullCommand():
		createLicense("boml", fileName, "", "")
	}
}
