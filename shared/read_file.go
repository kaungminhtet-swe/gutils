package shared

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func OpenFile(filepath string) (*os.File, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0444)

	if errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("file does not exist")
	}

	if err != nil {
		return nil, fmt.Errorf("unknown error %s", err)
	}

	return file, nil
}
