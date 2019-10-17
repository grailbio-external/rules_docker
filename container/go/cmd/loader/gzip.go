package main

import (
	"compress/gzip"
	"io"
	"os"
)

// gzReadCloser is a wrapper around gzip.Reader which closes the underlying
// reader of gzip.Reader on Close.
type gzReadCloser struct {
	// Compressed gzip.Reader
	gr io.ReadCloser
	// Underlying reader
	r io.ReadCloser
}

func newGZReadCloser(path string) (*gzReadCloser, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		_ = r.Close()
		return nil, err
	}
	return &gzReadCloser{gr: gr, r: r}, nil
}

func (gzr gzReadCloser) Read(p []byte) (int, error) {
	return gzr.gr.Read(p)
}

func (gzr gzReadCloser) Close() error {
	if err := gzr.r.Close(); err != nil {
		return err
	}
	return gzr.gr.Close()
}
