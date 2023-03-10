package fields

import (
	"example.com/utils"
	"github.com/jdkato/prose/v2"
)

var VoterIdLabel = "pii-indian-voter-id"

type VoterIdEntry = utils.ProdigyOutput

var sampleTextVoterId = []string{
	"voter %v",
	"voterid %v",
	"votercard %v",
	"voteridcard %v",
	"electoral photo identity card %v",
	"EPIC %v",
	"ECI %v",
	"election commmision %v",
}

var training_statements_VoterId = []utils.ProdigyOutput{
	VoterIdEntry{
		Text: sampleTextVoterId[0],
		Spans: []prose.LabeledEntity{
			{Start: 5, End: utils.GetWholeEntity(sampleTextVoterId[0], 5), Label: CSINLabel},
		},
	},
	VoterIdEntry{
		Text: sampleTextVoterId[1],
		Spans: []prose.LabeledEntity{
			{Start: 8, End: utils.GetWholeEntity(sampleTextVoterId[1], 8), Label: CSINLabel},
		},
	},

	VoterIdEntry{
		Text: sampleTextVoterId[2],
		Spans: []prose.LabeledEntity{
			{Start: 10, End: utils.GetWholeEntity(sampleTextVoterId[2], 10), Label: CSINLabel},
		},
	},
	VoterIdEntry{
		Text: sampleTextVoterId[3],
		Spans: []prose.LabeledEntity{
			{Start: 12, End: utils.GetWholeEntity(sampleTextVoterId[3], 12), Label: CSINLabel},
		},
	},
	VoterIdEntry{
		Text: sampleTextVoterId[4],
		Spans: []prose.LabeledEntity{
			{Start: 30, End: utils.GetWholeEntity(sampleTextVoterId[4], 30), Label: CSINLabel},
		},
	},
	VoterIdEntry{
		Text: sampleTextVoterId[5],
		Spans: []prose.LabeledEntity{
			{Start: 5, End: utils.GetWholeEntity(sampleTextVoterId[5], 5), Label: CSINLabel},
		},
	},

	VoterIdEntry{
		Text: sampleTextVoterId[6],
		Spans: []prose.LabeledEntity{
			{Start: 4, End: utils.GetWholeEntity(sampleTextVoterId[6], 4), Label: CSINLabel},
		},
	},
	VoterIdEntry{
		Text: sampleTextVoterId[7],
		Spans: []prose.LabeledEntity{
			{Start: 20, End: utils.GetWholeEntity(sampleTextVoterId[7], 20), Label: CSINLabel},
		},
	},
}

func GetTextForFieldVoterId(voterId string) VoterIdEntry {
	return utils.GetRandomEntry(training_statements_VoterId, voterId)
}
