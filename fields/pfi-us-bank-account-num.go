package fields

import (
	"example.com/utils"
)

var USBankAccountLabel = "pfi-us-bank-account-number"

type USBankAccountEntry = utils.ProdigyOutput

var sampleTextUSBankAccountLabel = []string{
	"Checking Account Number %v",
	"Checking Account %v",
	"Checking Account # %v",
	"Checking Acct Number %v",
	"Checking Acct # %v",
	"Checking Acct No. %v",
	"Checking Account No. %v",
	"Bank Account Number %v",
	"Bank Account # %v",
	"Bank Acct Number %v",
	"Bank Acct # %v",
	"Bank Acct No. %v",
	"Bank Account No. %v",
	"Savings Account Number %v",
	"Savings Account. %v",
	"Savings Account # %v",
	"Savings Acct Number %v",
	"Savings Acct # %v",
	"Savings Acct No. %v",
	"Savings Account No. %v",
	"Debit Account Number %v",
	"Debit Account %v",
	"Debit Account # %v",
	"Debit Acct Number %v",
	"Debit Acct # %v",
	"Debit Acct No. %v",
	"Debit Account No. %v",
}
