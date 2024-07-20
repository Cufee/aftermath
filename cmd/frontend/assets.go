package frontend

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/internal/log"
)

func calculateETag(content []byte) string {
	h := sha256.New()
	h.Write(content)
	hash := hex.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("\"%s\"", hash)
}

func NewAssetsHandler(root fs.FS) handler.Servable {
	// Generate ETags for all assets
	var tags = make(map[string]string)
	err := fs.WalkDir(root, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := fs.ReadFile(root, path)
		if err != nil {
			log.Fatal().Err(err).Str("path", path).Msg("failed to read assets file")
		}
		tags["/assets/"+path] = calculateETag(data)
		return err
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read assets fs")
	}

	fs := http.StripPrefix("/assets/", http.FileServerFS(root))
	return handler.HTTP(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Delete if-modified-since header so that ETags can be used instead of the standard cache policy.
		r.Header.Del("If-Modified-Since")

		// Set ETag to to uniquely identify the unchanged static asset.
		w.Header().Set("ETag", tags[r.URL.Path])

		fs.ServeHTTP(w, r)
	}))
}
