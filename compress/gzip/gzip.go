package gzip

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/kivigo/encoders/model"
)

var _ model.Compressor = (*compressor)(nil)

func New() model.Compressor {
	return compressor{}
}

// compressor provides methods to compress and decompress data using the gzip algorithm.
type compressor struct{}

// Compress compresses the input data using gzip and returns the compressed bytes.
func (compressor) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(data)
	w.Close()
	return buf.Bytes(), err
}

// Decompress decompresses gzip-compressed data and returns the original bytes.
func (compressor) Decompress(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

// Name returns the name of the compressor ("gzip").
func (compressor) Name() string { return "gzip" }
