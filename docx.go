package go_docx

import "github.com/google/uuid"

type DOCX struct {
	tmpUnzippedStructure string
	unzipped             bool
	uuid                 uuid.UUID
	mediaFiles           []string
}

func New() *DOCX {
	return &DOCX{
		uuid: uuid.New(),
	}
}
