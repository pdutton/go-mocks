#!/bin/make

GO?=go
MOCKGEN?=mockgen

.PHONY: all
all: \
  net/http/mock_http/http.go \
  net/http/mock_http/client.go \
  net/http/mock_http/client_request.go \
  net/http/mock_http/client_response.go \
  net/http/mock_http/server.go

.PHONY: mockgen
mockgen: 
	$(GO) install go.uber.org/mock/mockgen@latest

net/http/mock_http/http.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http HTTP

net/http/mock_http/client.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http Client

net/http/mock_http/client_request.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http ClientRequest

net/http/mock_http/client_response.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http ClientResponse

net/http/mock_http/server.go:
	$(MOCKGEN) -destination $@ github.com/pdutton/go-interfaces/net/http Server

