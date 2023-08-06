package go_docx

func (d *DOCX) GetMediaFilesName() ([]string, error) {
	if !d.unzipped { // if not unzipped we cant read this one
		return nil, ErrFileHasNotBeenReadYet
	}

	return d.mediaFiles, nil
}
