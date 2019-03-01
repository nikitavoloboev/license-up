package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// cc0Create creates the CC0 1.0 Universal license
func cc0Create(fileName string) error {
	fo, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	CC0 := "CC0 1.0 Universal\n\n" +
		"Statement of Purpose\n\n" +
		"The laws of most jurisdictions throughout the world automatically confer\n" +
		"exclusive Copyright and Related Rights (defined below) upon the creator and\n" +
		"subsequent owner(s) (each and all, an \"owner\") of an original work of\n" +
		"authorship and/or a database (each, a \"Work\").\n\n" +
		"Certain owners wish to permanently relinquish those rights to a Work for the\n" +
		"purpose of contributing to a commons of creative, cultural and scientific\n" +
		"works (\"Commons\") that the public can reliably and without fear of later\n" +
		"claims of infringement build upon, modify, incorporate in other works, reuse\n" +
		"and redistribute as freely as possible in any form whatsoever and for any\n" +
		"purposes, including without limitation commercial purposes. These owners may\n" +
		"contribute to the Commons to promote the ideal of a free culture and the\n" +
		"further production of creative, cultural and scientific works, or to gain\n" +
		"reputation or greater distribution for their Work in part through the use and\n" +
		"efforts of others.\n\n" +
		"For these and/or other purposes and motivations, and without any expectation\n" +
		"of additional consideration or compensation, the person associating CC0 with a\n" +
		"Work (the \"Affirmer\"), to the extent that he or she is an owner of Copyright\n" +
		"and Related Rights in the Work, voluntarily elects to apply CC0 to the Work" +
		"and publicly distribute the Work under its terms, with knowledge of his or her\n" +
		"Copyright and Related Rights in the Work and the meaning and intended legal\n" +
		"effect of CC0 on those rights.\n\n" +
		"1. Copyright and Related Rights. A Work made available under CC0 may be\n" +
		"protected by copyright and related or neighboring rights (\"Copyright and\n" +
		"Related Rights\"). Copyright and Related Rights include, but are not limited\n" +
		"to, the following:\n\n" +
		"  i. the right to reproduce, adapt, distribute, perform, display, communicate,\n" +
		"  and translate a Work;\n\n" +
		"  ii. moral rights retained by the original author(s) and/or performer(s);\n\n" +
		"  iii. publicity and privacy rights pertaining to a person's image or likeness\n" +
		"  depicted in a Work;\n\n" +
		"  iv. rights protecting against unfair competition in regards to a Work,\n" +
		"  subject to the limitations in paragraph 4(a), below;\n\n" +
		"  v. rights protecting the extraction, dissemination, use and reuse of data in\n" +
		"  a Work;\n\n" +
		"  vi. database rights (such as those arising under Directive 96/9/EC of the\n" +
		"  European Parliament and of the Council of 11 March 1996 on the legal\n" +
		"  protection of databases, and under any national implementation thereof,\n" +
		"  including any amended or successor version of such directive); and\n\n" +
		"  vii. other similar, equivalent or corresponding rights throughout the world\n" +
		"  based on applicable law or treaty, and any national implementations thereof.\n\n" +
		"2. Waiver. To the greatest extent permitted by, but not in contravention of,\n" +
		"applicable law, Affirmer hereby overtly, fully, permanently, irrevocably and\n" +
		"unconditionally waives, abandons, and surrenders all of Affirmer's Copyright\n" +
		"and Related Rights and associated claims and causes of action, whether now\n" +
		"known or unknown (including existing as well as future claims and causes of\n" +
		"action), in the Work (i) in all territories worldwide, (ii) for the maximum\n" +
		"duration provided by applicable law or treaty (including future time\n" +
		"extensions), (iii) in any current or future medium and for any number of\n" +
		"copies, and (iv) for any purpose whatsoever, including without limitation\n" +
		"commercial, advertising or promotional purposes (the \"Waiver\"). Affirmer makes\n" +
		"the Waiver for the benefit of each member of the public at large and to the\n" +
		"detriment of Affirmer's heirs and successors, fully intending that such Waiver\n" +
		"shall not be subject to revocation, rescission, cancellation, termination, or\n" +
		"any other legal or equitable action to disrupt the quiet enjoyment of the Work\n" +
		"by the public as contemplated by Affirmer's express Statement of Purpose.\n\n" +
		"3. Public License Fallback. Should any part of the Waiver for any reason be\n" +
		"judged legally invalid or ineffective under applicable law, then the Waiver\n" +
		"shall be preserved to the maximum extent permitted taking into account\n" +
		"Affirmer's express Statement of Purpose. In addition, to the extent the Waiver\n" +
		"is so judged Affirmer hereby grants to each affected person a royalty-free,\n" +
		"non transferable, non sublicensable, non exclusive, irrevocable and\n" +
		"unconditional license to exercise Affirmer's Copyright and Related Rights in\n" +
		"the Work (i) in all territories worldwide, (ii) for the maximum duration\n" +
		"provided by applicable law or treaty (including future time extensions), (iii)\n" +
		"in any current or future medium and for any number of copies, and (iv) for any\n" +
		"purpose whatsoever, including without limitation commercial, advertising or\n" +
		"promotional purposes (the \"License\"). The License shall be deemed effective as\n" +
		"of the date CC0 was applied by Affirmer to the Work. Should any part of the\n" +
		"License for any reason be judged legally invalid or ineffective under\n" +
		"applicable law, such partial invalidity or ineffectiveness shall not\n" +
		"invalidate the remainder of the License, and in such case Affirmer hereby\n" +
		"affirms that he or she will not (i) exercise any of his or her remaining\n" +
		"Copyright and Related Rights in the Work or (ii) assert any associated claims\n" +
		"and causes of action with respect to the Work, in either case contrary to\n" +
		"Affirmer's express Statement of Purpose.\n\n" +
		"4. Limitations and Disclaimers.\n\n" +
		"  a. No trademark or patent rights held by Affirmer are waived, abandoned,\n" +
		"  surrendered, licensed or otherwise affected by this document.\n\n" +
		"  b. Affirmer offers the Work as-is and makes no representations or warranties\n" +
		"  of any kind concerning the Work, express, implied, statutory or otherwise,\n" +
		"  including without limitation warranties of title, merchantability, fitness\n" +
		"  for a particular purpose, non infringement, or the absence of latent or\n" +
		"  other defects, accuracy, or the present or absence of errors, whether or not\n" +
		"  discoverable, all to the greatest extent permissible under applicable law.\n\n" +
		"  c. Affirmer disclaims responsibility for clearing rights of other persons\n" +
		"  that may apply to the Work or any use thereof, including without limitation\n" +
		"  any person's Copyright and Related Rights in the Work. Further, Affirmer\n" +
		"  disclaims responsibility for obtaining any necessary consents, permissions\n" +
		"  or other rights required for any use of the Work.\n\n" +
		"  d. Affirmer understands and acknowledges that Creative Commons is not a\n" +
		"  party to this document and has no duty or obligation with respect to this\n" +
		"  CC0 or use of the Work.\n\n" +
		"For more information, please see\n" +
		"<http://creativecommons.org/publicdomain/zero/1.0/>\n"

	ioutil.WriteFile(fileName, []byte(CC0), 0644)
	fmt.Println("License was created")
	return nil
}
