package main

import (
	"strings"
	"unicode"

	"github.com/otiai10/gosseract/v2"
)

var instance *gosseract.Client

// TextMethod contains configuration of tesseract client
type TextMethod struct {
	Name      string
	Language  string
	Variables map[string]string
	Client    *gosseract.Client
}

func (tm TextMethod) extract(data []byte) (*string, error) {

	err := tm.Client.SetImageFromBytes(data)
	if err != nil {
		return nil, err
	}

	output, err := tm.Client.Text()
	if len(strings.TrimSpace(output)) == 0 || err != nil {
		return nil, err
	}

	result := strings.TrimFunc(output, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	return &result, nil
}

func (tm TextMethod) tesseractSettings() {
	tm.Client.SetLanguage(tm.Language)

	// tesseract parameters:
	// http://www.sk-spell.sk.cx/tesseract-ocr-parameters-in-302-version
	for k, v := range tm.Variables {
		tm.Client.SetVariable(gosseract.SettableVariable(k), v)
	}

}
