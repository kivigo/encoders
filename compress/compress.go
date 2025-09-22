package compress

import (
	"context"

	"github.com/kivigo/encoders/model"
)

var _ model.Encoder = (*compressEncoder)(nil)

type compressEncoder struct {
	encoder    model.Encoder
	compressor model.Compressor
}

func New(encoder model.Encoder, compressor model.Compressor) model.Encoder {
	return &compressEncoder{
		encoder:    encoder,
		compressor: compressor,
	}
}

func (c *compressEncoder) Encode(ctx context.Context, value any) ([]byte, error) {
	data, err := c.encoder.Encode(ctx, value)
	if err != nil {
		return nil, err
	}
	return c.compressor.Compress(data)
}

func (c *compressEncoder) Decode(ctx context.Context, data []byte, value any) error {
	raw, err := c.compressor.Decompress(data)
	if err != nil {
		return err
	}
	return c.encoder.Decode(ctx, raw, value)
}
