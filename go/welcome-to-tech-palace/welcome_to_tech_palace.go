package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return string("Welcome to the Tech Palace," + " " + strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// var a = oldMsg
	// var b = strings.Replace(a, "*", "", -1)
	// var c = strings.TrimSpace(b)
	// return c
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
