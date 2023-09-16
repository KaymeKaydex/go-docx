package go_docx

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (d *DOCX) FromFile(path string) (*DOCX, error) {
	if path == "" {
		return nil, fmt.Errorf("path is empty")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = d.describeAsArchive(data)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *DOCX) FromReader(reader io.Reader) (*DOCX, error) {
	if reader == nil {
		return nil, fmt.Errorf("nil reader pointer")
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = d.describeAsArchive(data)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *DOCX) describeAsArchive(data []byte) error {
	err := os.WriteFile(fmt.Sprintf("tmp/%s.zip", d.uuid), data, 0777)
	if err != nil {
		return err
	}

	err = d.unzip()
	if err != nil {
		return err
	}

	err = d.rmArchive()
	if err != nil {
		return err
	}

	d.unzipped = true

	return nil
}

func (d *DOCX) rmArchive() error {
	return os.Remove(fmt.Sprintf("tmp/%s.zip", d.uuid))
}

func (d *DOCX) unzip() error {
	archive, err := zip.OpenReader(fmt.Sprintf("%s/%s.zip", d.tmpDir, d.uuid))
	if err != nil {
		return err
	}

	defer archive.Close()

	if archive == nil {
		return fmt.Errorf("arch is nil")
	}

	err = os.Mkdir(fmt.Sprintf("tmp/%s", d.uuid), 0777)
	if err != nil {
		return err
	}

	mediaFiles := make([]string, 0)

	dst := fmt.Sprintf("tmp/%s", d.uuid)
	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		if strings.HasPrefix(f.Name, "word/media/") { // there are media type as jpg/png and other
			mediaFiles = append(mediaFiles, strings.TrimPrefix(f.Name, "word/media/"))
		}

		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path")
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	d.mediaFiles = mediaFiles

	return nil
}

func (d *DOCX) rmUnarchivedFolder() error {
	return os.RemoveAll(fmt.Sprintf("tmp/%s", d.uuid))
}

func (d *DOCX) Close() error {
	d.unzipped = false
	return d.rmUnarchivedFolder()
}
