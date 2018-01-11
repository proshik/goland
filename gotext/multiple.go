package main

import (
	"fmt"
	"github.com/leonelquinteros/gotext"
)

func main() {
	// Create Locale with library path and language code
	ru := gotext.NewLocale("./gotext/locales", "ru_RU")

	// Load domain './gotext/locales/default.po'
	ru.AddDomain("default")

	// Translate text from default domain
	fmt.Println(ru.Get("lamb"))

	// Load different domain
	ru.AddDomain("translations")

	// Translate text from domain
	fmt.Println(ru.GetD("translations", "lamb"))

	// en_US
	en := gotext.NewLocale("./gotext/locales", "en_US")

	fmt.Println(en.Get("lamb"))

}
