package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// unlicenseCreate creates the Unlicense license
func unlicenseCreate(fileName string) error {
	fo, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	Unlicense := "This is free and unencumbered software released into the public domain.\n\n" +
		"Anyone is free to copy, modify, publish, use, compile, sell, or\n" +
		"distribute this software, either in source code form or as a compiled\n" +
		"binary, for any purpose, commercial or non-commercial, and by any\n" +
		"means.\n\n" +
		"In jurisdictions that recognize copyright laws, the author or authors\n" +
		"of this software dedicate any and all copyright interest in the\n" +
		"software to the public domain. We make this dedication for the benefit\n" +
		"of the public at large and to the detriment of our heirs and\n" +
		"successors. We intend this dedication to be an overt act of\n" +
		"relinquishment in perpetuity of all present and future rights to this\n" +
		"software under copyright law.\n\n" +
		"THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND,\n" +
		"EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF\n" +
		"MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.\n" +
		"IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR\n" +
		"OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,\n" +
		"ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR\n" +
		"OTHER DEALINGS IN THE SOFTWARE.\n\n" +
		"For more information, please refer to <http://unlicense.org>\n"

	ioutil.WriteFile(fileName, []byte(Unlicense), 0644)
	fmt.Println("License was created")
	return nil
}
