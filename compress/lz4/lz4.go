package lz4

import (
	"bytes"
	"io"

	"github.com/kivigo/encoders/model"
	"github.com/pierrec/lz4/v4"
)

var _ model.Compressor = (*compressor)(nil)

func New() model.Compressor {
	return compressor{}
}

// compressor provides methods to compress and decompress data using the lz4 algorithm.
type compressor struct{}

// Compress compresses the input data using lz4 and returns the compressed bytes.
func (compressor) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := lz4.NewWriter(&buf)
	_, err := w.Write(data)
	if err != nil {
		w.Close()
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress decompresses lz4-compressed data and returns the original bytes.
func (compressor) Decompress(data []byte) ([]byte, error) {
	r := lz4.NewReader(bytes.NewReader(data))
	return io.ReadAll(r)
}

// Name returns the name of the compressor ("lz4").
func (compressor) Name() string { return "lz4" }
