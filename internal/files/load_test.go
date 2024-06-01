package files

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFiles(t *testing.T) {
	mockFs := fstest.MapFS{
		"file.txt": {
			Data: []byte("file.txt value"),
		},
		"sub_dir/file.txt": {
			Data: []byte("sub_dir/file.txt value"),
		},
	}

	files, err := ReadDirFiles(mockFs, ".")
	assert.NoError(t, err)
	assert.Equal(t, string(files["file.txt"]), "file.txt value")
	assert.Equal(t, string(files["sub_dir/file.txt"]), "sub_dir/file.txt value")

	_, err = ReadDirFiles(nil, ".")
	assert.ErrorIs(t, err, fs.ErrInvalid)
}
