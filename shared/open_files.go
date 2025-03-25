package shared

import (
	"log/slog"
	"os"
	"path/filepath"
)

func OpenFiles(filepaths ...string) ([]*os.File, error) {
	files := make([]*os.File, 0)

	for _, path := range filepaths {
		file, err := OpenFile(path)
		if err != nil {
			slog.Error("file", filepath.Base(path), err.Error())
			continue
		}
		files = append(files, file)
	}

	return files, nil
}
