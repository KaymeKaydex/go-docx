package go_docx

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type DOCX struct {
	tmpUnzippedStructure string
	unzipped             bool
	uuid                 uuid.UUID
	tmpDir               string
	mediaFiles           []string
}

func New() *DOCX {
	return &DOCX{
		uuid:   uuid.New(),
		tmpDir: "tmp",
	}
}

func (d *DOCX) GetContentTypeXML() (XMLContentType, error) {
	if !d.unzipped {
		return XMLContentType{}, ErrFileHasNotBeenReadYet
	}

	bts, err := os.ReadFile(fmt.Sprintf("tmp/%s/[Content_Types].xml", d.uuid))
	if err != nil {
		return XMLContentType{}, err
	}

	ct := &XMLContentType{}

	err = xml.Unmarshal(bts, ct)
	if err != nil {
		return XMLContentType{}, err
	}

	return *ct, err
}

type WordStructure struct {
	xmlContentType XMLContentType
	docPropsFolder docPropsFolder
	wordFolder     wordFolder
}

type XMLContentType struct {
	XMLName xml.Name `xml:"Types"`

	Overrides []XMLContentTypeOverride `xml:"Override"`
	Defaults  []XMLContentTypeDefault  `xml:"Default"`
}

type XMLContentTypeDefault struct {
	Extension   string `xml:"Extension,attr"`
	ContentType string `xml:"ContentType,attr"`
}

type XMLContentTypeOverride struct {
	PartName    string `xml:"PartName,attr"`
	ContentType string `xml:"ContentType,attr"`
}

type wordFolder struct {
}

type docPropsFolder struct {
}
