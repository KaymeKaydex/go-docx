package go_docx

import (
	"fmt"
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
	fmt.Println(ct)
}
