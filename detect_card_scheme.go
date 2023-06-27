package main

type CardScheme struct {
	Name     string
	Lengths  []LenRange
	Prefixes []PrefixRange
}

type LenRange struct {
	Min, Max int
}

type PrefixRange struct {
	Min, Max string
}

var CardSchemes = []CardScheme{
	{
		Name: "American Express",
		Lengths: []LenRange{
			{15, 15},
		},
		Prefixes: []PrefixRange{
			{"34", "34"},
			{"37", "37"},
		},
	},
	{
		Name: "JCB",
		Lengths: []LenRange{
			{16, 19},
		},
		Prefixes: []PrefixRange{
			{"3528", "3589"},
		},
	},
	{
		Name: "Maestro",
		Lengths: []LenRange{
			{12, 19},
		},
		Prefixes: []PrefixRange{
			{"50", "50"},
			{"56", "58"},
			{"6", "6"},
		},
	},
	{
		Name: "Visa",
		Lengths: []LenRange{
			{13, 13},
			{16, 16},
			{19, 19},
		},
		Prefixes: []PrefixRange{
			{"4", "4"},
		},
	},
	{
		Name: "MasterCard",
		Lengths: []LenRange{
			{16, 16},
		},
		Prefixes: []PrefixRange{
			{"2221", "2720"},
			{"51", "55"},
		},
	},
}

func DetectCardScheme(cardNumber string) string {
	cardNumber = sanitizeCardNumber(cardNumber)
	for _, scheme := range CardSchemes {
		if checkLength(cardNumber, scheme.Lengths) && checkRange(cardNumber, scheme.Prefixes) && IsValidCardNumber(cardNumber) {
			return scheme.Name
		}
	}

	return "Unknown"
}

func checkLength(cardNumber string, lengths []LenRange) bool {
	cardNumberLength := len(cardNumber)

	for _, length := range lengths {
		if cardNumberLength >= length.Min && cardNumberLength <= length.Max {
			return true
		}
	}

	return false
}

func checkRange(cardNumber string, ranges []PrefixRange) bool {
	for _, r := range ranges {
		if cardNumber[0:len(r.Min)] >= r.Min && cardNumber[0:len(r.Max)] <= r.Max {
			return true
		}
	}

	return false
}
