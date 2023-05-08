package fields

import (
	"errors"
	"regexp"
	"strconv"
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
	ValidateFn:    ValidateUsBankAccountNumber,
	Regex:         `\b\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d\s*\d{0,4}\b`,
	SensitiveInfo: []int{},
}

var UsItin = Field{
	Name:          "PFI_US_ITIN",
	ValidateFn:    ValidateUsITIN,
	Regex:         `(9\d{2}[-\s]?(5[0-9]|[6][0-5]|[7-8][0-9]|[9][024-9])[-\s]?\d{4}|\d{9})`,
	SensitiveInfo: []int{},
}

var CanadaSocSecInsu = Field{
	Name:          "PII_CAN_SOC_INSU",
	ValidateFn:    ValidateCanadaSIN,
	Regex:         `(?:\d{3}[-\s]?\d{3}[-\s]?\d{3}|\d{9})`,
	SensitiveInfo: []int{},
}

var IndianAadhar = Field{
	Name:          "PII_INDIAN_AADHAR",
	ValidateFn:    ValidateAadhaar, //  Verhoeff algorithm’s
	Regex:         `[2-9]\d{2}[-\s]?\d{4}[-\s]?\d{4}`,
	SensitiveInfo: []int{},
}

var UsDriverLicense = Field{
	Name:          "PII_US_DRIVER_LICENSE",
	ValidateFn:    nil,
	Regex:         `[2-9]\d{2}[-\s]?\d{4}[-\s]?\d{4}`, // depends on the state
	SensitiveInfo: []int{},
}

var UsSSN = Field{
	Name:          "PII_US_SSN",
	ValidateFn:    ValidateSSN,
	Regex:         `^\d{3}-\d{2}-\d{4}$`,
	SensitiveInfo: []int{},
}

var UkPassportNum = Field{
	Name:          "PII_UK_Passport_Num",
	ValidateFn:    nil,
	Regex:         `^[A-Z]{2}\d{7}$`,
	SensitiveInfo: []int{},
}

var UsPassportNum = Field{
	Name:          "PII_US_Passport_Num",
	ValidateFn:    nil,
	Regex:         `^[A-Z]\d{8}$`,
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

func ValidateAadhaar(num string) bool {
	if len(num) != 12 {
		return false
	}

	for _, c := range num {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return ValidateVerhoeff(num)
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

func ValidateSSN(input string) bool {
	if len(input) == 9 { // unformatted
		for _, c := range input {
			if !unicode.IsDigit(c) {
				return false
			}
		}
	} else {

		if len(input) != 11 { // formatted
			return false
		}

		for i, c := range input {
			if i == 3 || i == 6 {
				if !(c == '-' || c == ' ') {
					return false
				}
			} else {
				if !unicode.IsDigit(c) {
					return false
				}
			}
		}
	}
	return true
}

func ValidateUsITIN(itin string) bool {
	// Verify that the ITIN is a 9-digit number
	if len(itin) != 9 {
		return false
	}
	for _, c := range itin {
		if c < '0' || c > '9' {
			return false
		}
	}

	// Verify the first two digits
	if itin[0] != '9' || itin[1] < '7' || itin[1] > '9' {
		return false
	}

	// Calculate the check digit
	weights := [8]int{2, 7, 6, 5, 4, 3, 2, 1}
	total := 0
	for i, c := range itin[:8] {
		digit := int(c - '0')
		total += digit * weights[i]
	}
	checkDigit := (11 - (total % 11)) % 10

	// Verify the check digit
	if checkDigit != int(itin[8]-'0') {
		return false
	}

	return true
}

/*
Validating US Bank Account Number

Validating a US bank account number involves checking whether the account number conforms to a specific format and verifying the number using a checksum. Here's an algorithm to validate a US bank account number:

Verify that the account number is between 8 and 12 digits long.
Verify that the first two digits of the account number are not 00.
If the account number is 8 digits long, it should be formatted as XXXXXXXY, where Y is a check digit.
If the account number is 9 digits long, it should be formatted as XXXXXYYYC, where YYY is a check digit and C is a routing number checksum.
If the account number is 10 digits long, it should be formatted as XXXYYYZZZZ, where YYY is a check digit and ZZZZ is a routing number.
If the account number is 11 digits long, it should be formatted as YYXXXXXYYYC, where YY is the Federal Reserve routing symbol, XXXX is the Federal Reserve office identifier, YYY is a check digit, and C is a routing number checksum.
If the account number is 12 digits long, it should be formatted as YYXXXYYYZZZZ, where YY is the Federal Reserve routing symbol, XXX is the Federal Reserve office identifier, YYY is a check digit, and ZZZZ is a routing number.

*/

func ValidateUsBankAccountNumber(accountNumber string) bool {
	// Verify that the account number is between 8 and 12 digits long
	if len(accountNumber) < 8 || len(accountNumber) > 12 {
		return false
	}

	// Verify that the first two digits of the account number are not 00
	if accountNumber[:2] == "00" {
		return false
	}

	// Verify the format and checksum based on the length of the account number
	switch len(accountNumber) {
	case 8:
		// Format: XXXXXXXY
		checkDigit := int(accountNumber[7] - '0')
		return checkDigit == calculateCheckDigit(accountNumber[:7])
	case 9:
		// Format: XXXXXYYYC
		checkDigit := cast.ToInt(accountNumber[6:9])
		routingNumberChecksum := int(accountNumber[8] - '0')
		return checkDigit == calculateCheckDigit(accountNumber[:6]) &&
			routingNumberChecksum == calculateRoutingNumberChecksum(accountNumber[:8])
	case 10:
		// Format: XXXYYYZZZZ
		checkDigit := cast.ToInt(accountNumber[6:9])
		routingNumber, _ := strconv.Atoi(accountNumber[6:])
		return checkDigit == calculateCheckDigit(accountNumber[:6]) && routingNumber >= 1 &&
			routingNumber <= 9999
	case 11:
		// Format: YYXXXXXYYYC
		checkDigit := cast.ToInt(accountNumber[6:9])
		routingNumberChecksum := int(accountNumber[10] - '0')
		return checkDigit == calculateCheckDigit(accountNumber[:5]+accountNumber[9:10]) &&
			routingNumberChecksum == calculateRoutingNumberChecksum(accountNumber[:10])
	case 12:
		// Format: YYXXXYYYZZZZ
		checkDigit := cast.ToInt(accountNumber[6:9])
		routingNumber, _ := strconv.Atoi(accountNumber[6:])
		return checkDigit == calculateCheckDigit(accountNumber[:5]+accountNumber[9:]) &&
			routingNumber >= 1 &&
			routingNumber <= 9999
	default:
		return false
	}
}

func calculateCheckDigit(number string) int {
	// Calculate the check digit using the weighted modulus 10 algorithm
	weights := []int{3, 7, 1, 3, 7, 1, 3, 7}
	sum := 0
	for i, digit := range number {
		if i >= len(weights) {
			break
		}
		sum += int(digit-'0') * weights[i]
	}
	remainder := sum % 10
	if remainder == 0 {
		return 0
	} else {
		return 10 - remainder
	}
}

func calculateRoutingNumberChecksum(number string) int {
	// Calculate the routing number checksum using the weighted modulus 10 algorithm
	weights := []int{3, 7, 1, 3, 7, 1, 3}
	sum := 0
	for i, digit := range number {
		if i >= len(weights) {
			break
		}
		sum += int(digit-'0') * weights[i]
	}
	remainder := sum % 10
	return (10 - remainder) % 10
}

/*
Validating Canadian Social Security Number


Verify that the SIN is a string of nine digits.

Verify that the first digit is not "0" or "8".

Split the SIN into three groups of three digits each.

Verify that the first group is not "000".

Verify that the second group is not "00".

Use the following algorithm to calculate the checksum:

Multiply the first digit of the SIN by 1, the second digit by 2, the third digit by 1, the fourth digit by 2, and so on, alternating between multiplying by 1 and 2.

Sum the individual digits of the products obtained in step 1. If a product is greater than 9, sum the individual digits of that product. For example, if a digit is multiplied by 2 and the result is 14, sum 1 + 4 = 5.

Sum the results of step 2.

If the result of step 3 is a multiple of 10, the SIN is valid. Otherwise, it is invalid.
*/

func ValidateCanadaSIN(sin string) bool {
	if len(sin) != 9 || !regexp.MustCompile(`^\d{9}$`).MatchString(sin) {
		return false
	}

	if sin[0] == '0' || sin[0] == '8' {
		return false
	}

	groups := []string{sin[0:3], sin[3:6], sin[6:9]}

	if groups[0] == "000" || groups[1] == "00" {
		return false
	}

	weightedSum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(sin[i]))
		if i%2 == 0 {
			weightedSum += digit
		} else {
			product := digit * 2
			if product > 9 {
				weightedSum += product/10 + product%10
			} else {
				weightedSum += product
			}
		}
	}

	return weightedSum%10 == 0
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
