package wc

import (
	`fmt`
	"github.com/kaungminhtet-swe/gutils/shared"
	"github.com/stretchr/testify/assert"
	"io"
	`log/slog`
	"os"
	"path"
	`path/filepath`
	"strings"
	"testing"
)

var file *os.File

func TestMain(m *testing.M) {
	// Get data file path
	cd, _ := os.Getwd()
	splitcd := strings.Split(cd, "/")
	rootpath := strings.Join(splitcd[:len(splitcd)-1], "/")
	filepath := path.Join(rootpath, "test/data/test.txt")

	file, _ = shared.OpenFile(filepath)

	code := m.Run()

	if err := file.Close(); err != nil {
		slog.Error("Error found in closing file", file.Name(), err.Error())
	}

	os.Exit(code)
}

func Test_Count_Lines(t *testing.T) {
	testcases := []struct {
		name     string
		input    io.Reader
		expected int64
		err      string
	}{
		{
			name:     "null input",
			input:    nil,
			expected: 0,
			err:      "nil reader",
		},
		{
			name:     "empty line",
			input:    strings.NewReader(""),
			expected: 0,
			err:      "",
		},
		{
			name:     "one line",
			input:    strings.NewReader("Hello, World!"),
			expected: 1,
			err:      "",
		},
		{
			name: "two line",
			input: strings.NewReader(`Hello, World!
Hello, World!
Hello, World!`),
			expected: 3,
			err:      "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := countLines(tc.input)
			assert.Equal(t, tc.expected, actual)

			if tc.err == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tc.err)
			}
		})
	}
}

func TestCountLinesFromFile(t *testing.T) {
	actual := countLinesFromFiles([]*os.File{file})
	assert.Equal(t, []int64{7145}, actual)
}

func TestCountLines(t *testing.T) {
	cd, _ := os.Getwd()
	splitcd := strings.Split(cd, "/")
	rootpath := strings.Join(splitcd[:len(splitcd)-1], "/")
	testpath := path.Join(rootpath, "test/data/test.txt")

	lines := CountLines(testpath)

	assert.Equal(t, fmt.Sprintf("%d lines %s\n", 7145,
		filepath.Base(testpath)), lines)
}
