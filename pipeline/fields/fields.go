package fields

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/spf13/cast"
)

const (
	asciiZero = 48
	asciiTen  = 57
)

const (
	nilType = iota

	pii_age
	pfi_us_bank_account_num
	pii_bmi
	pii_canada_insu_num
	// pii_eu_driver_license
	pii_indian_aadhar
	pii_indian_voter_id
	pii_money
	pii_region
	pii_sex
	pfi_us_itin
	pfi_credit_card
	// pii_us_driver_license
	date
	name
)

type Field struct {
	Name          string            `json:"name"`
	ValidateFn    func(string) bool `json:"validation"`
	Regex         string            `json:"regex"`
	SensitiveInfo []int             `json:"sensitiveInfo"` // first element is the word range and rest are the dataTypes
}

var CreaditCard = Field{
	Name:          "PFI_CREDIT_CARD",
	ValidateFn:    ValidateCreditCard,
	Regex:         "",
	SensitiveInfo: []int{300, date, name},
}

var UsBankAccountNumber = Field{
	Name:          "PFI_USBAN",
	ValidateFn:    nil,
	Regex:         `\b\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d{0,4}\b`,
	SensitiveInfo: []int{},
}

var UsItin = Field{
	Name:          "PFI_US_ITIN",
	ValidateFn:    nil,
	Regex:         `(9\d{2}[-\s]?(5[0-9]|[6][0-5]|[7-8][0-9]|[9][024-9])[-\s]?\d{4}|\d{9})`,
	SensitiveInfo: []int{},
}

var CanadaSocSecInsu = Field{
	Name:          "PII_CAN_SOC_INSU",
	ValidateFn:    nil, // Luhn algorithm
	Regex:         `(?:\d{3}[-\s]?\d{3}[-\s]?\d{3}|\d{9})`,
	SensitiveInfo: []int{},
}

var IndianAadhar = Field{
	Name:          "PII_INDIAN_AADHAR",
	ValidateFn:    ValidateVerhoeff, //  Verhoeff algorithm’s
	Regex:         `[2-9]\d{2}[-\s]?\d{4}[-\s]?\d{4}`,
	SensitiveInfo: []int{},
}

var UsDriverLicense = Field{
	Name:          "PII_US_DRIVER_LICENSE",
	ValidateFn:    nil,
	Regex:         `[2-9]\d{2}[-\s]?\d{4}[-\s]?\d{4}`, // depends on the state
	SensitiveInfo: []int{},
}

func MatchRegex(expression, input string) bool {
	re := regexp.MustCompile(expression)
	return re.MatchString(input)
}

// Validate returns an error if the provided string does not pass the luhn check.
func ValidateCreditCard(number string) bool {
	p := len(number) % 2
	sum, err := calculateLuhnSum(number, p)
	if err != nil {
		// return err
		return false
	}

	// If the total modulo 10 is not equal to 0, then the number is invalid.
	if sum%10 != 0 {
		// return errors.New("invalid number")
		return false
	}

	// return nil
	return true
}

func calculateLuhnSum(number string, parity int) (int64, error) {
	var sum int64
	for i, d := range number {
		if d < asciiZero || d > asciiTen {
			return 0, errors.New("invalid digit")
		}

		d = d - asciiZero
		// Double the value of every second digit.
		if i%2 == parity {
			d *= 2
			// If the result of this doubling operation is greater than 9.
			if d > 9 {
				// The same final result can be found by subtracting 9 from that result.
				d -= 9
			}
		}

		// Take the sum of all the digits.
		sum += int64(d)
	}

	return sum, nil
}

// The multiplication table
var verhoeff_d = [][]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
	{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
	{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
	{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
	{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
	{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
	{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
	{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
}

// The permutation table
var verhoeff_p = [][]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
	{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
	{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
	{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
	{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
	{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
	{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
}

// ValidateVerhoeff returns true if the passed string 'num' is Verhoeff compliant.  The check digit must be the last one.
func ValidateVerhoeff(num string) bool {
	if !IsInt(num) {
		return false
	}
	c := 0
	ll := len(num)
	for i := 0; i < ll; i++ {
		c = verhoeff_d[c][verhoeff_p[(i % 8)][num[ll-i-1]-'0']]
	}
	return (c == 0)
}

func IsUnderAge(age string) bool {
	return cast.ToInt(age) < 18
}

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
