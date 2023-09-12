package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var MedicalRecordNumberLabel = "PHI_MRN"

type MedicalRecordNumberEntry = utils.ProdigyOutput

var sampleTextMedicalRecordNumber = []string{
	" medical record number %v ",
	" mrn %v ",
	" %v ",
}

var training_statements_medical_record_number = []MedicalRecordNumberEntry{
	{
		Text: sampleTextMedicalRecordNumber[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 23,
				End:   utils.GetWholeEntity(sampleTextMedicalRecordNumber[0], 23),
				Label: MedicalRecordNumberLabel,
			},
		},
	},
	{
		Text: sampleTextMedicalRecordNumber[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 5,
				End:   utils.GetWholeEntity(sampleTextMedicalRecordNumber[1], 5),
				Label: MedicalRecordNumberLabel,
			},
		},
	},
	{
		Text: sampleTextMedicalRecordNumber[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextMedicalRecordNumber[2], 1),
				Label: MedicalRecordNumberLabel,
			},
		},
	},
}

func GetTextForFieldMedcialRecordNumber(mrn string) MedicalRecordNumberEntry {
	return utils.GetRandomEntry(training_statements_medical_record_number, mrn)
}
