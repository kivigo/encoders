package encrypt

import (
	"context"
	"errors"
	"testing"

	"github.com/azrod/cryptio"
)

// mockEncoder is a simple mock for model.Encoder
type mockEncoder struct {
	encodeFunc func(ctx context.Context, value any) ([]byte, error)
	decodeFunc func(ctx context.Context, data []byte, value any) error
}

func (m *mockEncoder) Encode(ctx context.Context, value any) ([]byte, error) {
	return m.encodeFunc(ctx, value)
}

func (m *mockEncoder) Decode(ctx context.Context, data []byte, value any) error {
	return m.decodeFunc(ctx, data, value)
}

func TestNewDefaultValue(t *testing.T) {
	_, err := New("", &mockEncoder{}, cryptio.SecurityMedium, cryptio.ProfileBalanced)
	if err == nil {
		t.Error("expected error for empty passphrase")
	}

	_, err = New("pass", nil, cryptio.SecurityMedium, cryptio.ProfileBalanced)
	if err == nil {
		t.Error("expected error for nil encoder")
	}

	enc, err := New("pass", &mockEncoder{
		encodeFunc: func(ctx context.Context, value any) ([]byte, error) { return []byte("ok"), nil },
		decodeFunc: func(ctx context.Context, data []byte, value any) error { return nil },
	}, cryptio.SecurityMedium, cryptio.ProfileBalanced)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if enc == nil {
		t.Error("expected encoder instance")
	}
}

func TestEncodeDecode(t *testing.T) {
	ctx := context.Background()
	mock := &mockEncoder{
		encodeFunc: func(ctx context.Context, value any) ([]byte, error) {
			if value == "fail" {
				return nil, errors.New("encode error")
			}
			return []byte("mockdata"), nil
		},
		decodeFunc: func(ctx context.Context, data []byte, value any) error {
			if string(data) == "fail" {
				return errors.New("decode error")
			}
			ptr, ok := value.(*string)
			if ok {
				*ptr = "decoded"
			}
			return nil
		},
	}
	enc, err := New("pass", mock, cryptio.SecurityMedium, cryptio.ProfileBalanced)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	e := enc.(*Encoder)

	// Test Encode nil value
	b, err := e.Encode(ctx, nil)
	if err != nil || b != nil {
		t.Errorf("expected nil, got %v, err %v", b, err)
	}

	// Test Encode error from underlying encoder
	_, err = e.Encode(ctx, "fail")
	if err == nil {
		t.Error("expected encode error")
	}

	// Test Encode/Decode roundtrip
	val := "test"
	encData, err := e.Encode(ctx, val)
	if err != nil {
		t.Fatalf("unexpected encode error: %v", err)
	}
	var out string
	err = e.Decode(ctx, encData, &out)
	if err != nil {
		t.Fatalf("unexpected decode error: %v", err)
	}
	if out != "decoded" {
		t.Errorf("expected 'decoded', got %q", out)
	}

	// Test Decode with nil value
	err = e.Decode(ctx, encData, nil)
	if err == nil {
		t.Error("expected error for nil value")
	}

	// Test Decode with empty data
	err = e.Decode(ctx, []byte{}, &out)
	if err != nil {
		t.Errorf("expected nil error for empty data, got %v", err)
	}
}

// Optionally, test error from cryptio.New
func TestNew_CryptioError(t *testing.T) {
	// Patch cryptio.New to return error if possible (skipped here, as cryptio is external)
	// This is a placeholder for completeness.
}
