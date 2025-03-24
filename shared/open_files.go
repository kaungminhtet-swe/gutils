package shared

import (
	`log/slog`
	`os`
)

func OpenFiles(filepaths ...string) ([]*os.File, error) {
	files := make([]*os.File, len(filepaths))

	for i, filepath := range filepaths {
		file, err := OpenFile(filepath)
		if err != nil {
			slog.Error("file", filepath, err.Error())
		}
		files[i] = file
	}

	return files, nil
}
