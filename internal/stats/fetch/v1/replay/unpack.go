package replay

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/cufee/aftermath/internal/log"
)

func UnpackRemote(link string) (*UnpackedReplay, error) {
	resp, err := http.DefaultClient.Get(link)
	if err != nil {
		return nil, ErrInvalidReplayFile
	}
	defer resp.Body.Close()

	// Convert 10 MB to bytes
	const maxFileSize = 10 * 1024 * 1024

	// Check the Content-Length header
	contentLengthStr := resp.Header.Get("Content-Length")
	if contentLengthStr == "" {
		log.Warn().Msg("Content-Length header is missing on remote replay file")
		return nil, ErrInvalidReplayFile
	}

	contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		return nil, err
	}
	if contentLength > maxFileSize {
		return nil, ErrInvalidReplayFile
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return Unpack(bytes.NewReader(data), contentLength)
}

// World of Tanks Blitz replay.
//
// # Replay structure
//
// `*.wotbreplay` is a ZIP-archive containing:
//
// - `battle_results.dat`
// - `meta.json`
// - `data.wotreplay`

type UnpackedReplay struct {
	BattleResult battleResults `json:"results"`
	Meta         replayMeta    `json:"meta"`
}

func Unpack(file io.ReaderAt, size int64) (*UnpackedReplay, error) {
	archive, err := newZipFromReader(file, size)
	if err != nil {
		return nil, err
	}
	if len(archive.File) < 3 {
		return nil, ErrInvalidReplayFile
	}

	var data UnpackedReplay

	resultsDat, err := archive.Open("battle_results.dat")
	if err != nil {
		return nil, ErrInvalidReplayFile
	}
	result, err := decodeBattleResults(resultsDat)
	if err != nil {
		return nil, ErrInvalidReplayFile
	}
	data.BattleResult = *result

	meta, err := archive.Open("meta.json")
	if err != nil {
		return nil, ErrInvalidReplayFile
	}
	metaBytes, err := io.ReadAll(meta)
	if err != nil {
		return nil, err
	}

	return &data, json.Unmarshal(metaBytes, &data.Meta)
}

func newZipFromReader(file io.ReaderAt, size int64) (*zip.Reader, error) {
	reader, err := zip.NewReader(file, size)
	if err != nil {
		return nil, err
	}

	return reader, nil
}
