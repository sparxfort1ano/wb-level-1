// Package adapter contains an adapter for a service,
// which historically intercats only with XML.
package adapter

import (
	"fmt"
	"io"
	"log"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/brianvoe/gofakeit/v6"
)

type XMLDocument struct {
	xmlData []byte
}

func (doc *XMLDocument) GetRandomXML() (err error) {
	doc.xmlData, err = gofakeit.XML(&gofakeit.XMLOptions{
		Type:          "single",
		RootElement:   "xml",
		RecordElement: "record",
		RowCount:      2,
		Indent:        true,
		Fields: []gofakeit.Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: gofakeit.MapParams{"special": {"false"}}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to rcv xml: %w", err)
	}
	return nil
}

type XMLDocumentAdapter struct {
	oldDoc *XMLDocument
}

func NewXMLDocumentAdapter(doc *XMLDocument) *XMLDocumentAdapter {
	return &XMLDocumentAdapter{oldDoc: doc}
}

func (adapter *XMLDocumentAdapter) SendJSONData(w io.Writer) error {
	xml := strings.NewReader(string(adapter.oldDoc.xmlData))
	jsonData, err := xj.Convert(xml)
	if err != nil {
		return fmt.Errorf("failed to convert xml to json: %w", err)
	}

	log.Println("JSON data post")
	if _, err := w.Write(jsonData.Bytes()); err != nil {
		return fmt.Errorf("failed to write json: %w", err)
	}
	return nil
}
