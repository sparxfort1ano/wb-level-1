// Package dataservice is like an external package,
// which interacts only with XML.
package dataservice

import (
	"fmt"
	"io"
	"log"

	"github.com/tidwall/randjson"
)

type DataService interface {
	SendJSONData(w io.Writer) error
}

type JSONDocument struct {
	jsonData []byte
}

func (doc *JSONDocument) GetRandomJSON() {
	doc.jsonData = randjson.Make(1, nil)
}

func (doc *JSONDocument) SendJSONData(w io.Writer) error {
	log.Println("JSON data post")
	if _, err := w.Write(doc.jsonData); err != nil {
		return fmt.Errorf("failed to write json: %w", err)
	}

	return nil
}
