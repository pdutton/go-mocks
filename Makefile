#!/bin/make

GO?=go
MOCKGEN?=mockgen

.DEFAULT_GOAL: all
.PHONY: all
all: \
  encoding/json/mock_json/all.go \
  io/mock_io/all.go \
  io/fs/mock_fs/all.go \
  net/mock_net/all.go \
  net/http/client/mock_client/all.go \
  net/http/server/mock_server/all.go \
  os/mock_os/all.go \
  os/exec/mock_exec/all.go \
  os/signal/mock_signal/all.go \
  path/mock_path/all.go \
  path/filepath/mock_filepath/all.go \
  sync/mock_sync/all.go

.PHONY: mockgen
mockgen:
	$(GO) install go.uber.org/mock/mockgen@latest

# Package github.com/pdutton/encoding/json:

.PHONY: encoding/json/mock_json/all.go
encoding/json/mock_json/all.go:
	$(MOCKGEN) -destination $@ -package mock_json github.com/pdutton/go-interfaces/encoding/json JSON,Decoder,Encoder

# Package github.com/pdutton/io:

.PHONY: io/mock_io/all.go
io/mock_io/all.go:
	$(MOCKGEN) -destination $@ -package mock_io github.com/pdutton/go-interfaces/io IO,ByteScanner,ByteWriter,Closer,ReadCloser,ReadSeekCloser,ReadSeeker,ReadWriteCloser,ReadWriteSeeker,ReadWriter,Reader,ReaderAt,ReaderFrom,RuneReader,RuneScanner,Seeker,StringWriter,WriteCloser,WriteSeeker,Writer,WriterAt,WriterTo,LimitedReader,OffsetWriter,PipeReader,PipeWriter,SectionReader

.PHONY: io/fs/mock_fs/all.go
io/fs/mock_fs/all.go:
	$(MOCKGEN) -destination $@ -package mock_fs github.com/pdutton/go-interfaces/io/fs FileSystem,DirEntry,FS,File,FileInfo,FileMode,GlobFS,ReadDirFS,ReadDirFile,ReadFileFS,StatFS,SubFS

# Package github.com/pdutton/net:

.PHONY: net/mock_net/all.go
net/mock_net/all.go:
	$(MOCKGEN) -destination $@ -package mock_net github.com/pdutton/go-interfaces/net Addr,Conn,Dialer,IPConn,ListenConfig,Listener,Net,PacketConn,Resolver,TCPConn,TCPListener,UDPConn,UnixConn,UnixListener 

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

# Package github.com/pdutton/io/signal:

.PHONY: os/signal/mock_signal/all.go
os/signal/mock_signal/all.go:
	$(MOCKGEN) -destination $@ -package mock_signal github.com/pdutton/go-interfaces/os/signal Signal

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


