package main

import "strconv"

func IsValidCardNumber(cardNumber string) bool {
	cardNumber = sanitizeCardNumber(cardNumber)

	var (
		l            = len(cardNumber)
		sum          int
		shouldDouble bool
	)

	for i := l - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}

		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble
	}

	return sum > 0 && sum%10 == 0
}
