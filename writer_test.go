package go_docx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDOCX_SaveDOCXFile(t *testing.T) {
	file, err := New().FromFile("testdata/test1.docx")
	require.NotNil(t, file)
	require.NoError(t, err)

	err = file.ExportDOCXFile("tmp/output.docx")
	require.NoError(t, err)
}
