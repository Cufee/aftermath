package generate

import (
	_ "github.com/Kodeworks/golang-image-ico"

	_ "github.com/kolesa-team/go-webp/encoder"
	_ "github.com/kolesa-team/go-webp/webp"
)

//go:generate go run ./generate.go
