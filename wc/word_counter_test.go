package wc_test

import (
	"fmt"
	"github.com/kaungminhtet-swe/gutils/wc"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

var testFilePath string

func TestMain(m *testing.M) {
	cd, _ := os.Getwd()
	splitcd := strings.Split(cd, "/")
	rootpath := strings.Join(splitcd[:len(splitcd)-1], "/")
	testFilePath = path.Join(rootpath, "test/data/test.txt")

	code := m.Run()

	os.Exit(code)
}

func TestWordsCount(t *testing.T) {
	word := wc.Run(wc.Word, testFilePath)
	assert.Equal(t, fmt.Sprintf("%d words %s\n", 58164,
		filepath.Base(testFilePath)),
		word)
}

func TestLinesCount(t *testing.T) {
	line := wc.Run(wc.Line, testFilePath)
	assert.Equal(t, fmt.Sprintf("%d lines %s\n", 7145,
		filepath.Base(testFilePath)),
		line)
}
