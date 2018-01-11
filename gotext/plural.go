package main

import (
	"fmt"
	"github.com/leonelquinteros/gotext"
)

func main() {
	// Set PO content
	str := `
msgid ""
msgstr ""

# Header below
"Plural-Forms: nplurals=2; plural=(n != 1);\n"

msgid "Translate this"
msgstr "Translated text"

msgid "Another string"
msgstr ""

msgid "One with var: %s"
msgid_plural "Several with vars: %s"
msgstr[0] "This one is the singular: %s"
msgstr[1] "This one is the plural: %s"
`

	// Create Po object
	po := new(gotext.Po)
	po.Parse(str)

	fmt.Println(po.GetN("One with var: %s", "Several with vars: %s", 54, "v"))
	// "This one is the plural: Variable"
}