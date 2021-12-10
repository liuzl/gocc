package gocc

import (
	"io"
	"os"
)

type Loader interface {
	Open(filepath string) (io.ReadCloser, error)
}

type OSLoader struct {
}

func (OSLoader) Open(filepath string) (io.ReadCloser, error) {
	return os.Open(filepath)
}
