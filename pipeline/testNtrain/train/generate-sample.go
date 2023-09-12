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
	pii_eu_driver_license
	pii_indian_aadhar
	pii_indian_voter_id
	pii_money
	pii_region
	pii_sex
	pfi_us_itin
	pfi_credit_card
	pii_name
	pii_country
	pii_fax
	pii_phone_number
	pii_social_media
	pii_web_url
	pii_zip_code
	pii_address
	pii_email
	pii_us_driver_license
	phi_us_hic
	phi_mrn
	pii_us_ssn
	phi_us_health_benificiary
	pii_us_vin
	pii_serial_number
	pii_date
)

var fieldMap = map[fieldType]fn{
	pii_age:                   fields.GetTextForFieldAge,
	pfi_us_bank_account_num:   fields.GetTextForFieldUSBankAccount,
	pii_bmi:                   fields.GetTextForFieldBMI,
	pii_canada_insu_num:       fields.GetTextForFieldCSIN,
	pii_us_driver_license:     fields.GetTextForFieldUsDriverLicence,
	pii_indian_aadhar:         fields.GetTextForFieldAadhar,
	pii_indian_voter_id:       fields.GetTextForFieldVoterId,
	pii_money:                 fields.GetTextForFieldMoney,
	pii_region:                fields.GetTextForFieldRegion,
	pii_sex:                   fields.GetTextForFieldSex,
	pfi_us_itin:               fields.GetTextForFieldUSItin,
	pfi_credit_card:           fields.GetTextForFieldCreditCard,
	pii_name:                  fields.GetTextForFieldName,
	pii_country:               fields.GetTextForFieldCountry,
	pii_fax:                   fields.GetTextForFieldFax,
	pii_phone_number:          fields.GetTextForFieldPhoneNumber,
	pii_social_media:          fields.GetTextForFieldSocialMedia,
	pii_web_url:               fields.GetTextForFieldWebUrl,
	pii_zip_code:              fields.GetTextForFieldZipCode,
	pii_address:               fields.GetTextForFieldAddress,
	pii_email:                 fields.GetTextForFieldEmail,
	phi_us_hic:                fields.GetTextForFieldHCI,
	phi_mrn:                   fields.GetTextForFieldMedcialRecordNumber,
	pii_us_ssn:                fields.GetTextForFieldUSSsn,
	phi_us_health_benificiary: fields.GetTextForFieldUSHealthBenificiaryNumber,
	pii_us_vin:                fields.GetTextForFieldUSVin,
	pii_serial_number:         fields.GetTextForFieldSerialNumber,
	pii_eu_driver_license:     fields.GetTextForFieldEuDriverLicence,
	pii_date:                  fields.GetTextForFieldDate,
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
	"walmart_stores.csv": {
		nilType,
		pii_name,
		pii_web_url,
		pii_address,
		pii_region,
		pii_region,
		pii_zip_code,
		pii_country,
		pii_phone_number,
		pii_phone_number,
		pii_fax,
		pii_fax,
		pii_email,
		pii_email,
		pii_web_url,
		nilType,
		nilType,
		nilType,
		pii_social_media,
		pii_social_media,
		pii_social_media,
		pii_social_media,
		pii_social_media,
	},
	"US-Health-insurance_claim-HIC-Number.csv": {phi_us_hic},
	"medical-record-number.csv":                {phi_mrn},
	"us-ssn.csv":                               {pii_us_ssn},
	"health-benificiary-number.csv":            {phi_us_health_benificiary},
	"us-driver-licence.csv":                    {pii_us_driver_license},
	"us-vin.csv":                               {pii_us_vin},
	"serial-number.csv":                        {pii_serial_number},
	"US-bank-account-number.csv": {
		pfi_us_bank_account_num,
		pii_date,
		nilType,
		nilType,
		pii_date,
		pii_money,
		pii_money,
		pii_money,
		nilType,
	},
}

// Account No,DATE,TRANSACTION DETAILS,CHQ.NO.,VALUE DATE, WITHDRAWAL AMT , DEPOSIT AMT ,BALANCE AMT,
func LoadRawData() {
	var inputText string

	root := "./dataset"

	files := utils.GetChildFiles(root)

	pathForTestFile := filepath.Join("./output/json/testFile.jsonl")
	for _, fileName := range files {
		fmt.Printf("Loading the file : %v\n\n", fileName.Name())
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
