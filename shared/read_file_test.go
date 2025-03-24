package shared

import (
	`github.com/stretchr/testify/assert`
	`os`
	`path`
	`strings`
	`testing`
)

func TestReadFile(t *testing.T) {
	t.Run("invalid file path", func(t *testing.T) {
		file, err := OpenFile("./test/data/test.txt")
		assert.Nil(t, file)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "file does not exist")
	})

	t.Run("valid file path", func(t *testing.T) {
		// Getting project directory path
		cd, err := os.Getwd()
		assert.Nil(t, err, err)
		splitcd := strings.Split(cd, "/")
		rootpath := strings.Join(splitcd[:len(splitcd)-1], "/")

		filepath := path.Join(rootpath, "test/data/test.txt")
		assert.FileExists(t, filepath)

		file, err := OpenFile(filepath)
		assert.NotNil(t, file)
		assert.Nil(t, err)

		file.Close()
	})
}
