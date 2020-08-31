PKGS := ./...
# SRCDIRS := $(shell go list -f '{{.Dir}}' $(PKGS))
GO := go

# check: misspell unconvert vet
 
vet: 
	$(GO) vet $(PKGS)

staticcheck:
	$(GO) get honnef.co/go/tools/cmd/staticcheck
	staticcheck -checks all $(PKGS)

misspell:
	$(GO) get github.com/client9/misspell/cmd/misspell
	misspell \
		-locale US \
		.

unconvert:
	$(GO) get github.com/mdempsky/unconvert
	unconvert -v $(PKGS)

ineffassign:
	$(GO) get github.com/gordonklaus/ineffassign
	find . -name '*.go' | xargs ineffassign

unparam:
	$(GO) get mvdan.cc/unparam
	unparam ./...

errcheck:
	$(GO) get github.com/kisielk/errcheck
	errcheck $(PKGS)
