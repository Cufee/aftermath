package files

import (
	"io/fs"

	"github.com/rs/zerolog/log"
)

func ReadDirFiles(dir fs.FS, path string) (map[string][]byte, error) {
	if dir == nil {
		return nil, fs.ErrInvalid
	}

	files := make(map[string][]byte)

	fs.WalkDir(dir, path, func(path string, d fs.DirEntry, err error) error {
		// If we are not able to access a file/directory for some reason, continue
		if err != nil {
			log.Err(err).Str("path", path).Msg("failed to read a file")
			return nil
		}

		if d.IsDir() {
			return nil
		}

		data, err := fs.ReadFile(dir, path)
		if err != nil {
			return err
		}
		files[path] = data
		return nil
	})

	return files, nil
}
