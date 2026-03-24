package main

import (
	"fmt"
	"os"

	dataservice "github.com/sparxfort1ano/wb-level-1/task21/data-service"
	adapter "github.com/sparxfort1ano/wb-level-1/task21/xml-adapter"
)

func main() {
	services := []dataservice.DataService{}

	// 1) JSON.
	doc1 := &dataservice.JSONDocument{}
	doc1.GetRandomJSON()
	services = append(services, doc1)

	// 2) XML adapter.
	xml := &adapter.XMLDocument{}
	if err := xml.GetRandomXML(); err != nil {
		fmt.Println(err.Error())
		return
	}
	doc2 := adapter.NewXMLDocumentAdapter(xml)
	services = append(services, doc2)

	for _, s := range services {
		if err := s.SendJSONData(os.Stdout); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
