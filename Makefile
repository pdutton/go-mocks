#!/bin/make

GO?=go
MOCKGEN?=mockgen

.PHONY: all
all: \
  io/mock_io/all.go \
  net/http/client/mock_client/all.go \
  net/http/server/mock_server/all.go \
  os/mock_os/all.go \
  sync/mock_sync/all.go

.PHONY: mockgen
mockgen: 
	$(GO) install go.uber.org/mock/mockgen@latest

# Package io:

.PHONY: io/mock_io/all.go
io/mock_io/all.go:
	$(MOCKGEN) -destination $@ io ByteScanner,ByteWriter,Closer,ReadCloser,ReadSeekCloser,ReadSeeker,ReadWriteCloser,ReadWriteSeeker,ReadWriter,Reader,ReaderAt,ReaderFrom,RuneReader,RuneScanner,Seeker,StringWriter,WriteCloser,WriteSeeker,Writer,WriterAt,WriterTo

# Package github.com/pdutton/net/http/client:

.PHONY: net/http/client/mock_client/all.go
net/http/client/mock_client/all.go:
	$(MOCKGEN) -destination $@ -package mock_http github.com/pdutton/go-interfaces/net/http/client HTTP,Client,Request,Response

# Package github.com/pdutton/net/http/server:

.PHONY: net/http/server/mock_server/all.go
net/http/server/mock_server/all.go:
	$(MOCKGEN) -destination $@ -package mock_http github.com/pdutton/go-interfaces/net/http/server HTTP,Server

# Package github.com/pdutton/os:

.PHONY: os/mock_os/all.go
os/mock_os/all.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/os File,FileInfo,OS,Process,Root

.PHONY: sync/mock_sync/all.go
sync/mock_sync/all.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/sync Cond,Locker,Map,Mutex,Once,Pool,RWMutex,WaitGroup


