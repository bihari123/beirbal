package train

type (
	fileName  = string
	fieldType = int
)

const (
	nil = iota

	pii_age
	pfi_us_bank_account_num
	pii_bmi
	pii_canada_insu_num
	pii_eu_driver_license
	pii_indian_aadhar
	pii_indian_voter_id
	pii_money
	pii_region
	pii_sex
	pii_us_driver_license
)

var csvFilemap = map[string][]fieldType{
	"insurance.csv": {pii_age, pii_sex, pii_bmi, nil, nil, pii_region, pii_money},
}

func loadRawData() {
}
