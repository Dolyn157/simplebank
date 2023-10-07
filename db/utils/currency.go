package utils

// Constants for currencies
const (
	USD = "USD"
	EUR = "EUR"
	JPY = "JPY"
	CAD = "CAD"
)

// isValidCurrency checks if the currency is valid
func IsValidCurrency(currency string) bool {
	switch currency {
	case USD, EUR, JPY, CAD:
		return true
	}
	return false
}
