package wc

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/kaungminhtet-swe/gutils/shared"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	Line = iota
	Word
)

func Run(flag int, filepaths ...string) string {
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
		}
	}()

	switch flag {
	case Line:
		return countLinesFromFiles(files...)
	case Word:
		return countWordsFromFiles(files...)
	default:
		return ""
	}
}

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

func countLinesFromFiles(files ...*os.File) string {
	var results = ""

	for i, file := range files {
		lines, err := countLines(file)
		if err != nil {
			slog.Error("Skipped: ", file.Name(), err.Error())
			continue
		}
		results += fmt.Sprintf("%d lines %s\n", lines,
			filepath.Base(files[i].Name()))
	}

	return results
}

func countWords(rd io.Reader) (int64, error) {
	scanner := bufio.NewScanner(rd)
	scanner.Split(bufio.ScanWords)

	var wordCount int64
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

func countWordsFromFiles(files ...*os.File) string {
	results := ""
	for _, file := range files {
		words, err := countWords(file)
		if err != nil {
			slog.Error("Skipped: ", file.Name(), err.Error())
			continue
		}
		results += fmt.Sprintf("%d words %s\n", words,
			filepath.Base(file.Name()))
	}

	return results
}
