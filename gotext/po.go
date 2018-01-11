package main

import (
	"fmt"
	"github.com/leonelquinteros/gotext"
)

func main() {
	// Set PO content
	str := `
msgid "Translate this"
msgstr "Translated text"

msgid "Another string"
msgstr ""

msgid "One with var: %s"
msgstr "This one sets the var: %s"
`

	// Create Po object
	po := new(gotext.Po)
	po.Parse(str)

	fmt.Println(po.Get("Translate this"))

	fmt.Println(po.Get("One with var: %s", "example"))
}
