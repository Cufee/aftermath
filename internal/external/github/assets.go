package github

import (
	"context"
	"io"
	"net/http"

	"github.com/google/go-github/v63/github"
	"github.com/pkg/errors"
)

var client = github.NewClient(nil)

func GetLatestGameAssets(ctx context.Context) (io.ReadCloser, int64, error) {
	release, _, err := client.Repositories.GetLatestRelease(ctx, "cufee", "aftermath-assets")
	if err != nil {
		return nil, 0, errors.Wrap(err, "gailed to get latest release")
	}

	for _, asset := range release.Assets {
		if *asset.Name == "assets.zip" {
			rc, _, err := client.Repositories.DownloadReleaseAsset(ctx, "cufee", "aftermath-assets", *asset.ID, http.DefaultClient)
			if err != nil {
				return nil, 0, errors.Wrap(err, "failed to download assets.zip")
			}

			return rc, int64(asset.GetSize()), nil
		}
	}
	return nil, 0, errors.New("assets not found in latest release")
}
