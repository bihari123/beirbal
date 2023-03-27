package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var USBankAccountLabel = "PFI-USBAN" // pfi-us-bank-account-number

type USBankAccountEntry = utils.ProdigyOutput

var sampleTextUSBankAccount = []string{
	"Checking Account Number %v ",
	"Checking Account %v ",
	"Checking Account # %v ",
	"Checking Acct Number %v ",
	"Checking Acct # %v ",
	"Checking Acct No. %v ",
	"Checking Account No. %v ",
	"Bank Account Number %v ",
	"Bank Account # %v ",
	"Bank Acct Number %v ",
	"Bank Acct # %v ",
	"Bank Acct No. %v ",
	"Bank Account No. %v ",
	"Savings Account Number %v ",
	"Savings Account. %v ",
	"Savings Account # %v ",
	"Savings Acct Number %v ",
	"Savings Acct # %v ",
	"Savings Acct No. %v ",
	"Savings Account No. %v ",
	"Debit Account Number %v ",
	"Debit Account %v ",
	"Debit Account # %v ",
	"Debit Acct Number %v ",
	"Debit Acct # %v ",
	"Debit Acct No. %v ",
	"Debit Account No. %v ",
}

var training_statements_us_bank_acc = []USBankAccountEntry{
	{
		Text: sampleTextUSBankAccount[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 24,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[0], 24),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 17,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[1], 17),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 19,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[2], 19),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[3],
		Spans: []prose.LabeledEntity{
			{
				Start: 21,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[3], 21),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[4],
		Spans: []prose.LabeledEntity{
			{
				Start: 16,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[4], 16),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[5],
		Spans: []prose.LabeledEntity{
			{
				Start: 18,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[5], 18),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[6],
		Spans: []prose.LabeledEntity{
			{
				Start: 21,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[6], 21),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[7],
		Spans: []prose.LabeledEntity{
			{
				Start: 20,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[7], 20),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[8],
		Spans: []prose.LabeledEntity{
			{
				Start: 15,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[8], 15),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[9],
		Spans: []prose.LabeledEntity{
			{
				Start: 17,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[9], 17),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[10],
		Spans: []prose.LabeledEntity{
			{
				Start: 12,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[10], 12),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[11],
		Spans: []prose.LabeledEntity{
			{
				Start: 14,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[11], 14),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[12],
		Spans: []prose.LabeledEntity{
			{
				Start: 17,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[12], 17),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[13],
		Spans: []prose.LabeledEntity{
			{
				Start: 23,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[13], 23),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[14],
		Spans: []prose.LabeledEntity{
			{
				Start: 17,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[14], 17),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[15],
		Spans: []prose.LabeledEntity{
			{
				Start: 18,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[15], 18),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[16],
		Spans: []prose.LabeledEntity{
			{
				Start: 20,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[16], 20),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[17],
		Spans: []prose.LabeledEntity{
			{
				Start: 15,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[17], 15),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[18],
		Spans: []prose.LabeledEntity{
			{
				Start: 17,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[18], 17),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[19],
		Spans: []prose.LabeledEntity{
			{
				Start: 20,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[19], 20),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[20],
		Spans: []prose.LabeledEntity{
			{
				Start: 21,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[20], 21),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[21],
		Spans: []prose.LabeledEntity{
			{
				Start: 14,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[21], 14),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[22],
		Spans: []prose.LabeledEntity{
			{
				Start: 16,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[22], 16),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[23],
		Spans: []prose.LabeledEntity{
			{
				Start: 18,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[23], 18),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[24],
		Spans: []prose.LabeledEntity{
			{
				Start: 13,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[24], 13),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[25],
		Spans: []prose.LabeledEntity{
			{
				Start: 15,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[25], 15),
				Label: USBankAccountLabel,
			},
		},
	},
	{
		Text: sampleTextUSBankAccount[26],
		Spans: []prose.LabeledEntity{
			{
				Start: 18,
				End:   utils.GetWholeEntity(sampleTextUSBankAccount[26], 18),
				Label: USBankAccountLabel,
			},
		},
	},
}

func GetTextForFieldUSBankAccount(accNumber string) USBankAccountEntry {
	return utils.GetRandomEntry(training_statements_us_bank_acc, accNumber)
}
