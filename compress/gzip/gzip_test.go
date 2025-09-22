package gzip

import (
	"bytes"
	"testing"
)

func TestGzipCompressor_CompressAndDecompress(t *testing.T) {
	comp := New()
	original := []byte("hello world, this is a test")

	// Test compression
	compressed, err := comp.Compress(original)
	if err != nil {
		t.Fatalf("Compress failed: %v", err)
	}
	if bytes.Equal(compressed, original) {
		t.Error("Compressed data should differ from original")
	}

	// Test decompression
	decompressed, err := comp.Decompress(compressed)
	if err != nil {
		t.Fatalf("Decompress failed: %v", err)
	}
	if !bytes.Equal(decompressed, original) {
		t.Errorf("Decompressed data mismatch: got %q, want %q", decompressed, original)
	}
}

func TestGzipCompressor_Decompress_InvalidData(t *testing.T) {
	comp := New()
	invalid := []byte("not a gzip stream")
	_, err := comp.Decompress(invalid)
	if err == nil {
		t.Error("Expected error when decompressing invalid data, got nil")
	}
}

func TestGzipCompressor_Name(t *testing.T) {
	comp := New()
	if comp.Name() != "gzip" {
		t.Errorf("Name() = %q, want %q", comp.Name(), "gzip")
	}
}
