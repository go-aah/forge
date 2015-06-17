package main

import (
	"fmt"

	"github.com/brettlangdon/forge"
)

func main() {
	// Parse a `SectionValue` from `example.cfg`
	settings, err := forge.ParseFile("example.cfg")
	if err != nil {
		panic(err)
	}

	// Get a single value
	if settings.Contains("global") {
		// Get `global` casted as `StringValue`
		value := settings.GetString("global")
		fmt.Printf("global = \"%s\"\r\n", value.GetValue())
	}

	// Get a nested value
	value, err := settings.Resolve("primary.included_setting")
	fmt.Printf("primary.included_setting = \"%s\"\r\n", value.GetValue())

	// Convert settings to a map
	settingsMap, err := settings.ToMap()
	fmt.Printf("global = \"%s\"\r\n", settingsMap["global"])

	// Convert settings to JSON
	jsonBytes, err := settings.ToJSON()
	fmt.Printf("\r\n\r\n%s\r\n", string(jsonBytes))
}
