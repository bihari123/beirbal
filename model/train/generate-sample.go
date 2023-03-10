package train

import (
	"path/filepath"

	"example.com/fields"
	"example.com/utils"
)

type (
	fileName  = string
	fieldType = int
	fn        func(string) string
)

const (
	nil = iota

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
	// pii_us_driver_license
)

var fieldMap = map[fieldType]fn{
	pii_age:                 fields.GetTextForFieldAge,
	pfi_us_bank_account_num: fields.GetTextForFieldUSBankAccount,
	pii_bmi:                 fields.GetTextForFieldBMI,
	pii_canada_insu_num:     fields.GetTextForFieldCSIN,
	// pii_eu_driver_license:fields.Get
	pii_indian_aadhar:   fields.GetTextForFieldAadhar,
	pii_indian_voter_id: fields.GetTextForFieldVoterId,
	pii_money:           fields.GetTextForFieldMoney,
	pii_region:          fields.GetTextForFieldRegion,
	pii_sex:             fields.GetTextForFieldSex,
}

var csvFilemap = map[string][]fieldType{
	"insurance": {pii_age, pii_sex, pii_bmi, nil, nil, pii_region, pii_money},
}

func loadRawData() {
	root := "./dataset"

	files := utils.GetChildFiles(root)

	for _, fileName := range files {
		if fields, ok := csvFilemap[fileName.Name()]; ok {
			fileNameWithExt := fileName.Name() + ".csv"
			data := utils.ReadCSV(filepath.Join(root, fileNameWithExt))

			for index, colData := range data {
				if index == 0 {
					continue
				}
			}

		} else {
			continue
		}
	}
}
