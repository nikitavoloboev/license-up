package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// bsd2Create creates the BSD 2-Clause license
func bsd2Create(name string, surname string) error {
	year, _, _ := time.Now().Date()
	fo, err := os.Create("LICENSE")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	BSD2 := "BSD 2-Clause License\n\n" +
		"Copyright (c) " + strconv.Itoa(year) + "-present, " + name + " " + surname + "\n" +
		"All rights reserved.\n\n" +
		"Redistribution and use in source and binary forms, with or without\n" +
		"modification, are permitted provided that the following conditions are met:\n\n" +
		"* Redistributions of source code must retain the above copyright notice, this\n" +
		"  list of conditions and the following disclaimer.\n\n" +
		"* Redistributions in binary form must reproduce the above copyright notice,\n" +
		"  this list of conditions and the following disclaimer in the documentation\n" +
		"  and/or other materials provided with the distribution.\n\n" +
		"THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\"\n" +
		"AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE\n" +
		"IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\n" +
		"DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\n" +
		"FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\n" +
		"DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\n" +
		"SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\n" +
		"CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\n" +
		"OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n" +
		"OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n"

	ioutil.WriteFile("LICENSE", []byte(BSD2), 0644)
	fmt.Println("License was created")
	return nil
}

// bsd3Create creates the BSD 3-Clause license
func bsd3Create(name string, surname string) error {
	year, _, _ := time.Now().Date()
	fo, err := os.Create("LICENSE")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	BSD3 := "BSD 3-Clause License\n\n" +
		"Copyright (c) " + strconv.Itoa(year) + "-present, " + name + " " + surname + "\n" +
		"All rights reserved.\n\n" +
		"Redistribution and use in source and binary forms, with or without\n" +
		"modification, are permitted provided that the following conditions are met:\n\n" +
		"* Redistributions of source code must retain the above copyright notice, this\n" +
		"  list of conditions and the following disclaimer.\n\n" +
		"* Redistributions in binary form must reproduce the above copyright notice,\n" +
		"  this list of conditions and the following disclaimer in the documentation\n" +
		"  and/or other materials provided with the distribution.\n\n" +
		"* Neither the name of the copyright holder nor the names of its\n" +
		"  contributors may be used to endorse or promote products derived from\n" +
		"  this software without specific prior written permission.\n\n" +
		"THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\"\n" +
		"AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE\n" +
		"IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\n" +
		"DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\n" +
		"FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\n" +
		"DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\n" +
		"SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\n" +
		"CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\n" +
		"OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n" +
		"OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n"

	ioutil.WriteFile("LICENSE", []byte(BSD3), 0644)
	fmt.Println("License was created")
	return nil
}
