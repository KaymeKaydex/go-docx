package go_docx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	file, err := os.Open("testdata/test1.docx")
	require.NoError(t, err) // check that all testdata inside library

	docx, err := New().FromReader(file)
	require.NoError(t, err)
	require.NotNil(t, docx)
}

func TestDOCX_FromFile(t *testing.T) {
	file, err := New().FromFile("testdata/test1.docx")
	require.NotNil(t, file)
	require.NoError(t, err)
	err = file.Close()
	require.NoError(t, err)
}
