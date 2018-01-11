package main

import (
	"fmt"
	"github.com/leonelquinteros/gotext"
)

func main() {
	// Configure package
	gotext.Configure("./gotext/locales", "ru_RU", "default")

	// Translate text from default domain
	fmt.Println(gotext.Get("Mary"))

	// Translate text from a different domain without reconfigure
	fmt.Println(gotext.GetD("domain2", "Another text on a different domain"))
}