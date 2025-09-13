package yaml

import (
	"context"
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/kivigo/encoders/model"
)

var _ model.Encoder = (*Encoder)(nil)

type Encoder struct{}

func New() model.Encoder {
	return &Encoder{}
}

// Encode encodes the given value into YAML format.
func (f *Encoder) Encode(ctx context.Context, value any) ([]byte, error) {
	if value == nil {
		return nil, nil
	}

	data, err := yaml.MarshalContext(ctx, value)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Decode decodes the given YAML data into the provided value.
func (f *Encoder) Decode(ctx context.Context, data []byte, value any) error {
	if value == nil {
		return fmt.Errorf("value cannot be nil")
	}

	if len(data) == 0 {
		return nil // No data to decode
	}

	err := yaml.UnmarshalContext(ctx, data, value)
	if err != nil {
		return err
	}

	return nil
}
