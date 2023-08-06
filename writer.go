package go_docx

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (d *DOCX) ExportDOCXFile(to string) error {
	if !d.unzipped {
		return ErrFileHasNotBeenReadYet
	}

	if !strings.HasSuffix(to, ".docx") {
		return fmt.Errorf("invalid file name; please choose correct like /tmp/my-file.docx")
	}

	file, err := os.Create(to)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", fmt.Sprintf("tmp/%s", d.uuid))
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.

		fPath := strings.TrimPrefix(path, fmt.Sprintf("tmp/%s/", d.uuid))
		f, err := w.Create(fPath)
		if err != nil {
			return err
		}

		fmt.Println(f)
		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	err = filepath.Walk(fmt.Sprintf("tmp/%s", d.uuid), walker)
	if err != nil {
		panic(err)
	}

	return nil
}
