package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
)

func create(name string, license string) error {
	fo, err := os.Create("LICENSE")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	MIT := "MIT License \n\n" +
		"Copyright (c) 2017 " + name + "\n\n" +
		"Permission is hereby granted, free of charge, to any person obtaining a copy \n" +
		"of this software and associated documentation files (the \"Software\"), to deal \n" +
		"in the Software without restriction, including without limitation the rights \n" +
		"to use, copy, modify, merge, publish, distribute, sublicense, and/or sell \n" +
		"copies of the Software, and to permit persons to whom the Software is \n" +
		"furnished to do so, subject to the following conditions: \n\n" +
		"The above copyright notice and this permission notice shall be included in all \n" +
		"copies or substantial portions of the Software.\n\n" +
		"THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n" +
		"IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n" +
		"FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n" +
		"AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n" +
		"LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n" +
		"OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE\n" +
		"SOFTWARE."

	ioutil.WriteFile("LICENSE", []byte(MIT), 0644)
	fmt.Println("License was created")
	return nil
}

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		name := c.Args().Get(0)
		name += " " + c.Args().Get(1)
		license := c.Args().Get(2)
		create(name, license)
		return nil
	}

	app.Run(os.Args)
}
