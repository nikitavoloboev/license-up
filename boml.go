package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// bomlCreate creates the Blue Oack Model license
func bomlCreate(fileName string, isMD bool) error {
	fo, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	mdBOML := "# Blue Oak Model License\n\n" +
		"Version 1.0.0\n\n" +
		"## Purpose\n\n" +
		"This license gives everyone as much permission to work with\n" +
		"this software as possible, while protecting contributors\n" +
		"from liability.\n\n" +
		"## Acceptance\n\n" +
		"In order to receive this license, you must agree to its\n" +
		"rules.  The rules of this license are both obligations\n" +
		"under that agreement and conditions to your license.\n" +
		"You must not do anything with this software that triggers\n" +
		"a rule that you cannot or will not follow.\n\n" +
		"## Copyright\n\n" +
		"Each contributor licenses you to do everything with this\n" +
		"software that would otherwise infringe that contributor's\n" +
		"copyright in it.\n\n" +
		"## Notices\n\n" +
		"You must ensure that everyone who gets a copy of\n" +
		"any part of this software from you, with or without\n" +
		"changes, also gets the text of this license or a link to\n" +
		"<https://blueoakcouncil.org/license/1.0.0>.\n\n" +
		"## Excuse\n\n" +
		"If anyone notifies you in writing that you have not\n" +
		"complied with [Notices](#notices), you can keep your\n" +
		"license by taking all practical steps to comply within 30\n" +
		"days after the notice.  If you do not do so, your license\n" +
		"ends immediately.\n\n" +
		"## Patent\n\n" +
		"Each contributor licenses you to do everything with this\n" +
		"software that would otherwise infringe any patent claims\n" +
		"they can license or become able to license.\n\n" +
		"## Reliability\n\n" +
		"No contributor can revoke this license.\n\n" +
		"## No Liability\n\n" +
		"***As far as the law allows, this software comes as is,\n" +
		"without any warranty or condition, and no contributor\n" +
		"will be liable to anyone for any damages related to this\n" +
		"software or this license, under any kind of legal claim.***\n"

	BOML := "Blue Oak Model License\n\n" +
		"Version 1.0.0\n\n" +
		"Purpose\n\n" +
		"This license gives everyone as much permission to work with\n" +
		"this software as possible, while protecting contributors\n" +
		"from liability.\n\n" +
		"Acceptance\n\n" +
		"In order to receive this license, you must agree to its\n" +
		"rules.  The rules of this license are both obligations\n" +
		"under that agreement and conditions to your license.\n" +
		"You must not do anything with this software that triggers\n" +
		"a rule that you cannot or will not follow.\n\n" +
		"Copyright\n\n" +
		"Each contributor licenses you to do everything with this\n" +
		"software that would otherwise infringe that contributor's\n" +
		"copyright in it.\n\n" +
		"Notices\n\n" +
		"You must ensure that everyone who gets a copy of\n" +
		"any part of this software from you, with or without\n" +
		"changes, also gets the text of this license or a link to\n" +
		"<https://blueoakcouncil.org/license/1.0.0>.\n\n" +
		"Excuse\n\n" +
		"If anyone notifies you in writing that you have not\n" +
		"complied with Notices, you can keep your\n" +
		"license by taking all practical steps to comply within 30\n" +
		"days after the notice.  If you do not do so, your license\n" +
		"ends immediately.\n\n" +
		"Patent\n\n" +
		"Each contributor licenses you to do everything with this\n" +
		"software that would otherwise infringe any patent claims\n" +
		"they can license or become able to license.\n\n" +
		"Reliability\n\n" +
		"No contributor can revoke this license.\n\n" +
		"No Liability\n\n" +
		"As far as the law allows, this software comes as is,\n" +
		"without any warranty or condition, and no contributor\n" +
		"will be liable to anyone for any damages related to this\n" +
		"software or this license, under any kind of legal claim.\n"

	if isMD {
		ioutil.WriteFile(fileName, []byte(mdBOML), 0644)
	} else {
		ioutil.WriteFile(fileName, []byte(BOML), 0644)
	}
	fmt.Println("License was created")
	return nil
}
