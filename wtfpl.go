package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// wtfplCreate creates the WTFPL license
func wtfplCreate(name string, surname string, fileName string) error {
	year, _, _ := time.Now().Date()
	fo, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	WTFPL := "        DO WHAT THE FUCK YOU WANT TO PUBLIC license\n" +
		"                    Version 2, December 2004\n\n" +
		"Copyright (c) " + strconv.Itoa(year) + "-present, " + name + " " + surname + "\n\n" +
		"Everyone is permitted to copy and distribute verbatim or modified\n" +
		"copies of this license document, and changing it is allowed as long\n" +
		"as the name is changed.\n\n" +
		"            DO WHAT THE FUCK YOU WANT TO PUBLIC license\n" +
		"   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION\n\n" +
		"0. You just DO WHAT THE FUCK YOU WANT TO.\n"

	ioutil.WriteFile(fileName, []byte(WTFPL), 0644)
	fmt.Println("License was created")
	return nil
}
