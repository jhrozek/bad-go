package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	// Register messages for different languages
	message.SetString(language.English, "price", "The price is %v.")
	message.SetString(language.Spanish, "price", "El precio es %v.")
	message.SetString(language.German, "price", "Der Preis beträgt %v.")

	// Define preferred languages
	preferredLanguages := []language.Tag{
		language.Spanish, // First preference
		language.German,  // Second preference
		language.English, // Third preference
	}

	// Match preferred language
	matcher := language.NewMatcher(preferredLanguages)
	lang, _, _ := matcher.Match(language.Spanish, language.English)

	// Print the price in the matched language
	p := message.NewPrinter(lang)
	price := 100.50
	p.Printf("price", price)

	// To demonstrate, let's print another message for English specifically
	p = message.NewPrinter(language.English)
	fmt.Println()
	p.Printf("price", price)
}
