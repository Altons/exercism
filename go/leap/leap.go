//Package leap provide methods to do calculation on Gregorian calendars
package leap

import "math"

// IsLeapYear return true is calendar year is a leap year. false otherwise
func IsLeapYear(year int) bool {
	y := float64(year)
	div4 := math.Mod(y, 4)
	div100 := math.Mod(y, 100)
	div400 := math.Mod(y, 400)

	if div4 == 0 {
		if div100 != 0 || div400 == 0 {
			return true
		}
	}
	return false
}
