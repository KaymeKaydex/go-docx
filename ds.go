package go_docx

import "encoding/xml"

type Document struct {
	XMLName xml.Name `xml:"w:document"`
	XMLW    string   `xml:"xmlns:w,attr"`
	XMLR    string   `xml:"xmlns:r,attr"`
	Body    *Body
}

type Body struct {
	XMLName   xml.Name `xml:"w:body"`
	Paragraph []*Paragraph
}

type File struct {
	Document    Document
	DocRelation DocRelation

	rId int
}

type Paragraph struct {
	XMLName xml.Name `xml:"w:p"`
	Data    []interface{}

	file *File
}

type DocRelation struct {
	XMLName      xml.Name        `xml:"Relationships"`
	XMLns        string          `xml:"xmlns,attr"`
	Relationship []*RelationShip `xml:"Relationship"`
}

type RelationShip struct {
	XMLName    xml.Name `xml:"Relationship"`
	ID         string   `xml:"Id,attr"`
	Type       string   `xml:"Type,attr"`
	Target     string   `xml:"Target,attr"`
	TargetMode string   `xml:"TargetMode,attr,omitempty"`
}
