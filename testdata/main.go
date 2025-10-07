package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/azrod/cryptio"
	"github.com/kivigo/encoders/encrypt"
)

// Simple JSON encoder
type jsonEncoder struct{}

func (j *jsonEncoder) Encode(ctx context.Context, value any) ([]byte, error) {
	return json.Marshal(value)
}
func (j *jsonEncoder) Decode(ctx context.Context, data []byte, value any) error {
	return json.Unmarshal(data, value)
}

// Utility function to compress with gzip
func compressGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	w.Close()
	return buf.Bytes(), nil
}

// Utility function to decompress with gzip
func decompressGzip(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

type BigStruct struct {
	Name    string
	Numbers []int
	Text    string
}

func main() {
	// Generate a large structure
	big := BigStruct{
		Name:    "Alice",
		Numbers: make([]int, 10000),
		Text:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. ",
	}

	fmt.Println("Large structure generation completed.")
	for i := range big.Numbers {
		big.Numbers[i] = i
	}
	fmt.Println("Numbers filling completed.")

	base := &jsonEncoder{}
	enc, err := encrypt.New("super-secret", base, cryptio.SecurityMedium, cryptio.ProfileBalanced)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// Encode raw JSON
	jsonRaw, err := base.Encode(ctx, big)
	if err != nil {
		panic(err)
	}

	// Encode + encrypt
	encrypted, err := enc.Encode(ctx, big)
	if err != nil {
		panic(err)
	}

	// Gzip compression on raw JSON
	jsonGz, err := compressGzip(jsonRaw)
	if err != nil {
		panic(err)
	}

	// Gzip compression on encrypted data
	encryptedGz, err := compressGzip(encrypted)
	if err != nil {
		panic(err)
	}

	// Gzip compression on raw JSON then encryption
	jsonGzEncrypted, err := enc.Encode(ctx, jsonGz)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Raw JSON size                : %d bytes\n", len(jsonRaw))
	fmt.Printf("Compressed JSON size         : %d bytes\n", len(jsonGz))
	fmt.Printf("Encrypted size               : %d bytes\n", len(encrypted))
	fmt.Printf("Encrypted compressed size    : %d bytes\n", len(encryptedGz))
	fmt.Printf("Compressed JSON encrypted size: %d bytes\n", len(jsonGzEncrypted))
}
