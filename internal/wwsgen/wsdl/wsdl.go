package wsdl

import "encoding/xml"

type Schema struct {
	XMLName       xml.Name    `xml:"definitions"`
	Documentation string      `xml:"documentation"`
	Messages      []Message   `xml:"message"`
	Operations    []Operation `xml:"portType>operation"`
	Elements      []Element   `xml:"types>schema>element"`
}

type Message struct {
	Name string      `xml:"name,attr"`
	Part MessagePart `xml:"part"`
}

type MessagePart struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

type Operation struct {
	Name          string         `xml:"name,attr"`
	Documentation string         `xml:"documentation"`
	Input         OperationInput `xml:"input"`
	Output        OperationInput `xml:"output"`
}

type OperationInput struct {
	Name    string `xml:"name,attr"`
	Message string `xml:"message,attr"`
}

type Element struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}
