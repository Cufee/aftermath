package common

import (
	"io"
)

type Image interface {
	PNG(io.Writer) error
}
