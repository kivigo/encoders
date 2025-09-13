<img align="left" width="140" src="https://raw.githubusercontent.com/kivigo/kivigo/refs/heads/main/website/static/img/logo-kivigo.png" alt="KiviGo Logo" />

# KiviGo — Encoders

Pluggable encoders for KiviGo. This small Go module provides a simple `Encoder` interface and ready-to-use implementations (JSON, YAML) so values can be serialized/deserialized consistently across KiviGo backends.

## Features

- Small and dependency-light encoder interface
- Built-in implementations:
  - JSON encoder (std lib)
  - YAML encoder (gopkg.in/yaml.v3)
- Easy to implement custom encoders (e.g. MsgPack, Gob, etc)
- Unit tested and CI-friendly

## Installation

```sh
go get github.com/kivigo/encoders/<encoder_name>
```

## Encoder interface

Implementations must satisfy a small interface used by the KiviGo client:

```go
package main

// Encoder defines how values are serialized and deserialized when stored or retrieved from a backend in KiviGo.
//
// KiviGo uses encoders to convert Go values (structs, strings, etc.) into a byte slice for storage,
// and to decode byte slices back into Go values when reading from the backend.
// This allows you to use different formats (JSON, YAML, etc.) or implement your own encoding logic.
//
// Example: using the JSON encoder, a struct will be marshaled to JSON before being saved in the database.
type Encoder interface {
 // Encode serializes the given value into a byte slice.
 //
 // Example:
 //   data, err := encoder.Encode("hello world")
 //   if err != nil {
 //       log.Fatal(err)
 //   }
 //   fmt.Println("Encoded:", data)
 Encode(v any) ([]byte, error)

 // Decode deserializes the given byte slice into the provided destination.
 //
 // Example:
 //   var s string
 //   err := encoder.Decode([]byte(`"hello world"`), &s)
 //   if err != nil {
 //       log.Fatal(err)
 //   }
 //   fmt.Println("Decoded:", s)
 Decode(data []byte, v any) error
}

```

## Usage with KiviGo

Use an encoder when creating the KiviGo client (example with JSON):

```go
import (
    "context"
    "github.com/azrod/kivigo"
    "github.com/azrod/kivigo/encoders/json"
    "github.com/azrod/kivigo/backend/local"
)

func example() error {
    bk, err := local.New(local.DefaultOptions("/tmp/data"))
    if err != nil { return err }

    c, err := kivigo.New(bk, kivigo.Option{Encoder: json.New()})
    if err != nil { return err }
    defer c.Close()

    if err := c.Set(context.Background(), "key", map[string]string{"foo":"bar"}); err != nil {
        return err
    }
    return nil
}
```

## Included encoders

- encoders/json — JSON encoder using `encoding/json`
- encoders/yaml — YAML encoder using `gopkg.in/yaml.v3`

## Writing a custom encoder

Create a type that implements `Encoder` and pass it to the client via `kivigo.Option{Encoder: myEncoder}`. Keep operations context-aware and error-returning.

## Tests

Run unit tests:

```sh
go test ./...
```

## Contributing

- Follow the project coding & testing guidelines
- Add unit tests for new encoders
- Document encoder-specific quirks in a short README or the project docs

## License

MIT — see the root repository LICENSE file.
