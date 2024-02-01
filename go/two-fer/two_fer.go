// Package twofer provides a guide for salespeople offering two-for-one deals.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns a response when a person asks for a cookie.
// If the name is empty, it returns a fallback.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
