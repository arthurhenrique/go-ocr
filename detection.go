package main

import (
	"strings"
	"sync"
	"unicode"

	"github.com/otiai10/gosseract/v2"
)

var once sync.Once
var instance *gosseract.Client

// TextMethod contains configuration of tesseract client
type TextMethod struct {
	Name      string
	Language  string
	Variables map[string]string
}

func (tm TextMethod) extract(data []byte) (*string, error) {
	client := tm.Gosseract()

	err := client.SetImageFromBytes(data)
	if err != nil {
		return nil, err
	}

	output, err := client.Text()
	if len(strings.TrimSpace(output)) == 0 || err != nil {
		return nil, err
	}

	result := strings.TrimFunc(output, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	return &result, nil
}

// Gosseract singleton implement
func (tm TextMethod) Gosseract() *gosseract.Client {
	once.Do(func() {
		instance = gosseract.NewClient()
		instance.SetLanguage(tm.Language)
		// tesseract parameters:
		// http://www.sk-spell.sk.cx/tesseract-ocr-parameters-in-302-version
		for k, v := range tm.Variables {
			instance.SetVariable(gosseract.SettableVariable(k), v)
		}
	})
	return instance
}
