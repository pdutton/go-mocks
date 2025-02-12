#!/bin/make

GO?=go
MOCKGEN?=mockgen

.PHONY: all
all: \
  io/mock_io/all.go \
  net/http/mock_http/all.go

.PHONY: mockgen
mockgen: 
	$(GO) install go.uber.org/mock/mockgen@latest

# Package io:

.PHONY: io/mock_io/all.go
io/mock_io/all.go:
	$(MOCKGEN) -destination $@ io ByteScanner,ByteWriter,Closer,ReadCloser,ReadSeekCloser,ReadSeeker,ReadWriteCloser,ReadWriteSeeker,ReadWriter,Reader,ReaderAt,ReaderFrom,RuneReader,RuneScanner,Seeker,StringWriter,WriteCloser,WriteSeeker,Writer,WriterAt,WriterTo

# Package github.com/pdutton/net/http:

.PHONY: net/http/mock_http/all.go
net/http/mock_http/all.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http HTTP,Client,ClientRequest,ClientResponse,Server


