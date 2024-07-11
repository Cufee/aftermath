package github

import (
	"context"
	"io"
	"testing"

	"github.com/matryer/is"
)

func TestGetLatestGameAssets(t *testing.T) {
	is := is.New(t)

	rc, size, err := GetLatestGameAssets(context.Background())
	is.NoErr(err)
	is.True(rc != nil)
	is.True(size > 0)
	defer rc.Close()

	data, err := io.ReadAll(rc)
	is.NoErr(err)
	is.True(data != nil)
}
