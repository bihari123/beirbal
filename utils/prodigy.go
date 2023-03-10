package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jdkato/prose/v2"
)

// ProdigyOutput represents a single entry of Prodigy's JSON Lines output.
//
// `LabeledEntity` is a structure defined by prose that specifies where the
// entities are within the given `Text`.
type ProdigyOutput struct {
	Text   string                `json:"text"`
	Spans  []prose.LabeledEntity `json:"spans"`
	Answer string                `json:"answer"`
}

/*
{"text":"This was taken during the Easter celebrations at Real de Catorce, MX.","spans":[{"start":66,"end":68,"label":"PRODUCT","answer":"reject"}]}
*/

func (p ProdigyOutput) ToString() string {
	return fmt.Sprintf(
		"{\"text\":\"%v\",\"spans\":[{\"start\":%v,\"end\":%v,\"label\":\"%v\"}],\"answer\":\"reject\"}",
		p.Text,
		p.Spans[0].Start,
		p.Spans[0].End,
		p.Spans[0].Label,
	)
}

// ReadProdigy reads our JSON Lines file line-by-line, populating a
// slice of `ProdigyOutput` structures.
func ReadProdigy(jsonLines []byte) []ProdigyOutput {
	dec := json.NewDecoder(bytes.NewReader(jsonLines))
	entries := []ProdigyOutput{}
	for {
		ent := ProdigyOutput{}
		err := dec.Decode(&ent)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		entries = append(entries, ent)
	}
	return entries
}

// Split divides our human-annotated data set into two groups: one for training
// our model and one for testing it.
//
// We're using an 80-20 split here, although you may want to use a different
// split.
func Split(data []ProdigyOutput) ([]prose.EntityContext, []ProdigyOutput) {
	cutoff := int(float64(len(data)) * 0.8)

	train, test := []prose.EntityContext{}, []ProdigyOutput{}
	for i, entry := range data {
		if i < cutoff {
			train = append(train, prose.EntityContext{
				Text:   entry.Text,
				Spans:  entry.Spans,
				Accept: entry.Answer == "accept",
			})
		} else {
			test = append(test, entry)
		}
	}

	return train, test
}
