#!/bin/make

GO?=go
MOCKGEN?=mockgen

.PHONY: all
all: \
  net/http/mock/http.go \
  net/http/mock/client.go \
  net/http/mock/client_request.go \
  net/http/mock/client_response.go \
  net/http/mock/server.go

.PHONY: mockgen
mockgen: 
	$(GO) install go.uber.org/mock/mockgen@latest

net/http/mock/http.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http HTTP

net/http/mock/client.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http Client

net/http/mock/client_request.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http ClientRequest

net/http/mock/client_response.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http ClientResponse

net/http/mock/server.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http Server

