package train

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/bihari123/beirbal/pipeline/fields"
	"github.com/bihari123/beirbal/pipeline/utils"
)

type (
	fileName  = string
	fieldType = int
	fn        func(string) utils.ProdigyOutput
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
	pfi_us_itin:         fields.GetTextForFieldUSItin,
	pfi_credit_card:     fields.GetTextForFieldCreditCard,
}

var csvFilemap = map[string][]fieldType{
	"insurance.csv": {
		pii_age,
		pii_sex,
		pii_bmi,
		nilType,
		nilType,
		pii_region,
		pii_money,
	},
	"Canada-Insu-Number.csv":   {pii_canada_insu_num},
	"Indian-Aadhar-Number.csv": {pii_indian_aadhar},
	"credit-card-number.csv":   {pfi_credit_card},
	"us-taxpayer-itin.csv":     {pfi_us_itin},
}

func LoadRawData() {
	var inputText string

	root := "./dataset"

	files := utils.GetChildFiles(root)

	pathForTestFile := filepath.Join("./output/json/testFile.jsonl")
	for _, fileName := range files {
		if fields, ok := csvFilemap[fileName.Name()]; ok {
			data := utils.ReadCSV(filepath.Join(root, fileName.Name()))

			for index, rowData := range data {
				if index == 0 {
					continue
				}

				for index2, colData := range rowData {
					fieldType := fields[index2]
					if fieldType == nilType {
						continue
					}

					inputText += fieldMap[fieldType](colData).ToString() + "\n"
				}
			}
			err := utils.WriteAppendFile(pathForTestFile, inputText)
			if err != nil {
				log.Fatal("error writing file", err)
			}
		} else {
			continue
		}
	}
	fmt.Println("file written at ", pathForTestFile)
}
