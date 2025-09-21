package encrypt

import (
	"context"
	"fmt"

	"github.com/azrod/cryptio"
	"github.com/kivigo/encoders/model"
)

var _ model.Encoder = (*Encoder)(nil)

type Encoder struct {
	passphrase string
	encoder    model.Encoder
	encrypter  *cryptio.Client
}

// New creates a new encrypting encoder wrapping the given encoder.
// More info about levels and profiles can be found in the (cryptio)[https://github.com/azrod/cryptio] documentation.
// Returns an error if the passphrase is empty or if the encoder is nil.
func New(passphrase string, encoder model.Encoder, level cryptio.SecurityLevel, profile cryptio.Argon2Profile) (model.Encoder, error) {
	if passphrase == "" {
		return nil, fmt.Errorf("passphrase cannot be empty")
	}

	if encoder == nil {
		return nil, fmt.Errorf("encoder cannot be nil")
	}

	encrypter, err := cryptio.New(passphrase, level, profile)
	if err != nil {
		return nil, fmt.Errorf("failed to create encrypter: %w", err)
	}

	return &Encoder{
		passphrase: passphrase,
		encoder:    encoder,
		encrypter:  encrypter,
	}, nil
}

// Encode encodes the given value using the underlying encoder and then encrypts the result.
func (f *Encoder) Encode(ctx context.Context, value any) ([]byte, error) {
	if value == nil {
		return nil, nil
	}

	// Encode the value using the underlying encoder
	data, err := f.encoder.Encode(ctx, value)
	if err != nil {
		return nil, err
	}

	// Encrypt the encoded data
	return f.encrypter.EncryptRaw(data)
}

// Decode decrypts the given data and delegates decoding to the underlying encoder.
func (f *Encoder) Decode(ctx context.Context, data []byte, value any) error {
	if value == nil {
		return fmt.Errorf("value cannot be nil")
	}

	if len(data) == 0 {
		return nil // No data to decode
	}

	// Decrypt the data
	dataDecrypted, err := f.encrypter.DecryptRaw(data)
	if err != nil {
		return err
	}

	// Decode the decrypted data using the underlying encoder
	return f.encoder.Decode(ctx, dataDecrypted, value)
}
