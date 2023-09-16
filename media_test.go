package go_docx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDOCX_GetMediaFilesName(t *testing.T) {
	file, err := New().FromFile("testdata/test1.docx")
	require.NotNil(t, file)
	require.NoError(t, err)

	media, err := file.GetMediaFilesName()
	require.NoError(t, err)
	require.NotNil(t, media)
	file.Close()
}

func TestDOCX_GetMediaFilesName2(t *testing.T) {
	file := New()

	media, err := file.GetMediaFilesName()
	require.Error(t, err)
	require.Nil(t, media)
}
