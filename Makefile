#!/bin/make

GO?=go
MOCKGEN?=mockgen

.PHONY: all
all: \
  io/mock_io/all.go \
  io/fs/mock_fs/all.go \
  net/http/client/mock_client/all.go \
  net/http/server/mock_server/all.go \
  os/mock_os/all.go \
  os/exec/mock_exec/all.go \
  path/mock_path/all.go \
  path/filepath/mock_filepath/all.go \
  sync/mock_sync/all.go

.PHONY: mockgen
mockgen: 
	$(GO) install go.uber.org/mock/mockgen@latest

# Package github.com/pdutton/io:

.PHONY: io/mock_io/all.go
io/mock_io/all.go:
	$(MOCKGEN) -destination $@ -package mock_io github.com/pdutton/go-interfaces/io IO,ByteScanner,ByteWriter,Closer,ReadCloser,ReadSeekCloser,ReadSeeker,ReadWriteCloser,ReadWriteSeeker,ReadWriter,Reader,ReaderAt,ReaderFrom,RuneReader,RuneScanner,Seeker,StringWriter,WriteCloser,WriteSeeker,Writer,WriterAt,WriterTo,LimitedReader,OffsetWriter,PipeReader,PipeWriter,SectionReader

.PHONY: io/fs/mock_fs/all.go
io/fs/mock_fs/all.go:
	$(MOCKGEN) -destination $@ -package mock_fs github.com/pdutton/go-interfaces/io/fs FileSystem,DirEntry,FS,File,FileInfo,FileMode,GlobFS,ReadDirFS,ReadDirFile,ReadFileFS,StatFS,SubFS

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

# Package github.com/pdutton/io/exec:

.PHONY: os/exec/mock_exec/all.go
os/exec/mock_exec/all.go:
	$(MOCKGEN) -destination $@ -package mock_exec github.com/pdutton/go-interfaces/os/exec Exec,Cmd

# Package path:

.PHONY: path/mock_path/all.go
path/mock_path/all.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/path Path

# Package path/filepath:

.PHONY: path/filepath/mock_filepath/all.go
path/filepath/mock_filepath/all.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/path/filepath DirEntry,FileInfo,FilePath

# Package sync:

.PHONY: sync/mock_sync/all.go
sync/mock_sync/all.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/sync Cond,Locker,Map,Mutex,Once,Pool,RWMutex,Sync,WaitGroup


