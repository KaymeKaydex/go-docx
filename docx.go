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

type Orientation int

const (
	OrientationVertical Orientation = iota
	OrientationHorizontal
)

func (d *DOCX) GetOrientation() (Orientation, error) {
	docInfo, err := d.GetWordDocumentXML()
	if err != nil {
		return 0, err
	}
	if docInfo.Body.Section.PgSz.Height > docInfo.Body.Section.PgSz.Width {
		return OrientationVertical, nil
	}

	return OrientationHorizontal, nil
}

// Full Word File

type WordStructure struct {
	xmlContentType XMLContentType
	docPropsFolder docPropsFolder
	wordFolder     wordFolder
}

type wordFolder struct {
}

type docPropsFolder struct {
}

// [Content_Types].xml

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

// word/document.xml

func (d *DOCX) GetWordDocumentXML() (XMLDocument, error) {
	if !d.unzipped {
		return XMLDocument{}, ErrFileHasNotBeenReadYet
	}

	bts, err := os.ReadFile(fmt.Sprintf("tmp/%s/word/document.xml", d.uuid))
	if err != nil {
		return XMLDocument{}, err
	}

	ct := &XMLDocument{}

	err = xml.Unmarshal(bts, ct)
	if err != nil {
		return XMLDocument{}, err
	}

	return *ct, err
}

type XMLDocument struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main document"`
	Body    Body     `xml:"body"`
}

type Body struct {
	XMLName    xml.Name    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Paragraphs []Paragraph `xml:"p"`
	Section    Section     `xml:"sectPr"`
}

type Paragraph struct {
	Text string `xml:"r>t"`
}

type Section struct {
	PgSz    PgSz    `xml:"pgSz"`
	PgMar   PgMar   `xml:"pgMar"`
	Cols    Cols    `xml:"cols"`
	DocGrid DocGrid `xml:"docGrid"`
}

type PgSz struct {
	Width  int `xml:"w,attr"`
	Height int `xml:"h,attr"`
}

type PgMar struct {
	Top    int `xml:"top,attr"`
	Right  int `xml:"right,attr"`
	Bottom int `xml:"bottom,attr"`
	Left   int `xml:"left,attr"`
	Header int `xml:"header,attr"`
	Footer int `xml:"footer,attr"`
	Gutter int `xml:"gutter,attr"`
}

type Cols struct {
	Space int `xml:"space,attr"`
}

type DocGrid struct {
	LinePitch int `xml:"linePitch,attr"`
}
