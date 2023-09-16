package go_docx

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDOCX_FromFile2(t *testing.T) {
	file, err := New().FromFile("testdata/test1.docx")
	require.NotNil(t, file)
	require.NoError(t, err)

	ct, err := file.GetContentTypeXML()
	require.NoError(t, err)
	require.NotNil(t, ct)
	file.Close()
}

func TestDOCX_GetContentWordDocumentXML(t *testing.T) {
	file, err := New().FromFile("testdata/test2.docx")
	require.NotNil(t, file)
	require.NoError(t, err)

	doc, err := file.GetWordDocumentXML()
	require.NoError(t, err)
	require.NotNil(t, doc)

	bts, err := xml.Marshal(doc)
	require.NoError(t, err)
	require.NotEmpty(t, bts)

	btsReal, err := os.ReadFile(fmt.Sprintf("tmp/%s/word/document.xml", file.uuid))
	require.NoError(t, err)

	require.Equal(t, string(btsReal), string(bts))

	file.Close()
}

func TestDOCX_GetOrientation(t *testing.T) {
	file, err := New().FromFile("testdata/test2.docx")
	require.NotNil(t, file)
	require.NoError(t, err)
	orientation, err := file.GetOrientation()
	require.NoError(t, err)
	require.Equal(t, orientation, OrientationVertical)
	file.Close()

	file, err = New().FromFile("testdata/test3.docx")
	require.NotNil(t, file)
	require.NoError(t, err)
	orientation, err = file.GetOrientation()
	require.NoError(t, err)
	require.Equal(t, orientation, OrientationHorizontal)
	file.Close()
}
