package wc_test

import (
	`github.com/kaungminhtet-swe/gutils/shared`
	`github.com/kaungminhtet-swe/gutils/wc`
	`github.com/stretchr/testify/assert`
	`io`
	`os`
	`path`
	`strings`
	"testing"
)

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
			actual, err := wc.CountLines(tc.input)
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
	// Get project root directory
	cd, err := os.Getwd()
	assert.Nil(t, err, err)
	splitcd := strings.Split(cd, "/")
	rootpath := strings.Join(splitcd[:len(splitcd)-1], "/")

	filepath := path.Join(rootpath, "test/data/test.txt")
	assert.FileExists(t, filepath)

	file, err := shared.ReadFile(filepath)
	assert.Nil(t, err)
	assert.NotNil(t, file)
	if file != nil {
		defer file.Close()
	}

	actual, err := wc.CountLines(file)
	assert.Nil(t, err)
	assert.Equal(t, int64(7145), actual)
}
