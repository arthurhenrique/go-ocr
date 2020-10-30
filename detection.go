package main

import (
	"strings"
	"unicode"

	"github.com/otiai10/gosseract/v2"
)

type TextMethod struct {
	Name      string
	Language  string
	Variables map[string]string
}

func (tm TextMethod) extract(data []byte) (*string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage(tm.Language)

	// tesseract parameters:
	// http://www.sk-spell.sk.cx/tesseract-ocr-parameters-in-302-version
	for k, v := range tm.Variables {
		client.SetVariable(gosseract.SettableVariable(k), v)
	}

	client.SetImageFromBytes(data)

	output, err := client.Text()
	if len(strings.TrimSpace(output)) == 0 || err != nil {
		return nil, err
	}

	result := strings.TrimFunc(output, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	return &result, nil
}
