package main

import "regexp"

func sanitizeCardNumber(cardNumber string) string {
	// Remove any non-digit characters from the card number
	re := regexp.MustCompile("[^0-9]")
	cardNumber = re.ReplaceAllString(cardNumber, "")

	return cardNumber
}
