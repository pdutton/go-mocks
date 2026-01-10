# go-mocks (Ha Ha!)

Mock implementations for Go standard library interfaces and those from [`github.com/pdutton/go-interfaces`](https://github.com/pdutton/go-interfaces) using [`go.uber.org/mock`](https://github.com/uber-go/mock).

## Installation

```bash
go get github.com/pdutton/go-mocks
```

## Usage

Import the mock package you need in your tests:

```go
import (
    "testing"
    mock_io "github.com/pdutton/go-mocks/io/mock_io"
    "go.uber.org/mock/gomock"
)

func TestYourFunction(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockReader := mock_io.NewMockReader(ctrl)
    mockReader.EXPECT().Read(gomock.Any()).Return(10, nil)

    // Use mockReader in your test
}
```

## Available Mocks

This repository provides mocks for the following packages:

- **encoding/json** - `JSON`, `Decoder`, `Encoder`
- **io** - `Reader`, `Writer`, `Closer`, `Seeker`, and more
- **io/fs** - `FS`, `File`, `DirEntry`, `FileInfo`, and more
- **net** - `Conn`, `Listener`, `Dialer`, `Resolver`, and more
- **net/http/client** - `Client`, `Request`, `Response`
- **net/http/server** - `Server`
- **os** - `File`, `FileInfo`, `Process`
- **os/exec** - `Cmd`, `Exec`
- **os/signal** - `Signal`
- **path** - `Path`
- **path/filepath** - `FilePath`, `DirEntry`, `FileInfo`
- **sync** - `Locker`, `Mutex`, `RWMutex`, `WaitGroup`, and more

## Generating Mocks

All mocks are auto-generated using `mockgen`. To regenerate:

```bash
# Install mockgen (if not already installed)
make mockgen

# Generate all mocks
make

# Generate specific package mock
make io/mock_io/all.go
make net/http/client/mock_client/all.go
```

## Adding New Mocks

To add mocks for a new package:

1. Add the target to the `all:` phony target in the Makefile
2. Create a new phony target following the pattern:
   ```makefile
   .PHONY: <path>/mock_<package>/all.go
   <path>/mock_<package>/all.go:
       $(MOCKGEN) -destination $@ -package mock_<package> github.com/pdutton/go-interfaces/<path> Interface1,Interface2,...
   ```
3. Run `make <path>/mock_<package>/all.go` to generate

## Dependencies

- Go 1.24.0+
- [`go.uber.org/mock`](https://github.com/uber-go/mock) - Mock generation framework
- [`github.com/pdutton/go-interfaces`](https://github.com/pdutton/go-interfaces) - Source interface definitions

## Important Notes

- **Never edit generated files**: All files in `mock_*/all.go` are auto-generated. Changes should be made to the Makefile and regenerated.
- **Version alignment**: When updating `go-interfaces` dependency, regenerate all mocks with `make all`

## License

See LICENSE file for details.
