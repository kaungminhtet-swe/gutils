package wc

import (
	"bufio"
	"errors"
	`fmt`
	"github.com/kaungminhtet-swe/gutils/shared"
	"io"
	"log"
	"log/slog"
	`os`
	`path/filepath`
)

func countLines(rd io.Reader) (int64, error) {
	if rd == nil {
		return 0, errors.New("nil reader")
	}

	br := bufio.NewReader(rd)
	var count int64
	for {
		_, _, err := br.ReadLine()
		if errors.Is(err, io.EOF) {
			return count, nil
		}

		if err != nil {
			return 0, err
		}

		count++
	}
}

func countLinesFromFiles(files []*os.File) []int64 {
	var results = make([]int64, len(files))

	for i, file := range files {
		lines, err := countLines(file)
		if err != nil {
			slog.Error("Skipped: ", file.Name(), err.Error())
			continue
		}
		results[i] = lines
	}

	return results
}

func CountLines(filepaths ...string) string {
	if len(filepaths) == 0 {
		log.Fatalf("empty files")
	}

	files, err := shared.OpenFiles(filepaths...)
	if err != nil || len(files) == 0 {
		log.Fatalf("can't read all files")
	}
	defer func() {
		for _, file := range files {
			if err := file.Close(); err != nil {
				slog.Error("Error found in closing file", file.Name(), err.Error())
			}
			slog.Info("File is closed", "file", filepath.Base(file.Name()))
		}
	}()

	lineCounts := countLinesFromFiles(files)

	var results string
	for i, lineCount := range lineCounts {
		results += fmt.Sprintf("%d lines %s\n", lineCount,
			filepath.Base(files[i].Name()))
	}

	return results
}
