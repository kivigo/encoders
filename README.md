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
package encoder

import "context"

type Encoder interface {
    // Encode transforms a Go value into bytes for storage.
    Encode(ctx context.Context, v any) ([]byte, error)

    // Decode transforms stored bytes back into the provided value pointer.
    Decode(ctx context.Context, data []byte, v any) error
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
