package json

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kivigo/encoders/model"
)

var _ model.Encoder = (*Encoder)(nil)

func New() model.Encoder {
	return &Encoder{}
}

type Encoder struct{}

// Encode encodes the given value into JSON format.
func (f *Encoder) Encode(_ context.Context, value any) ([]byte, error) {
	if value == nil {
		return nil, fmt.Errorf("value cannot be nil")
	}

	data, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("failed to encode value to JSON: %w", err)
	}

	return data, nil
}

// Decode decodes the given JSON data into the provided value.
func (f *Encoder) Decode(_ context.Context, data []byte, value any) error {
	if value == nil {
		return fmt.Errorf("value cannot be nil")
	}

	if len(data) == 0 {
		return nil // No data to decode
	}

	err := json.Unmarshal(data, value)
	if err != nil {
		return fmt.Errorf("failed to decode JSON data: %w", err)
	}

	return nil
}
