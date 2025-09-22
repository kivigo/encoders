package model

// Compressor defines an interface for data compression and decompression algorithms.
// It provides methods to compress and decompress byte slices, as well as a method to retrieve the compressor's name.
type Compressor interface {
	// Compress compresses the provided data and returns the result or an error if compression fails.
	Compress([]byte) ([]byte, error)
	// Decompress decompresses the provided data and returns the result or an error if decompression fails.
	Decompress([]byte) ([]byte, error)
	// Name returns the name of the compressor.
	Name() string
}
