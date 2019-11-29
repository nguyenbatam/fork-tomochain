.PHONY: fns fns-cross evm all test clean
.PHONY: fns-linux fns-linux-386 fns-linux-amd64 fns-linux-mips64 fns-linux-mips64le
.PHONY: fns-darwin fns-darwin-386 fns-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= latest
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

fns:
	build/env.sh go run build/ci.go install ./cmd/fns
	@echo "Done building."
	@echo "Run \"$(GOBIN)/fns\" to launch fns."

gc:
	build/env.sh go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	build/env.sh go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

fns-cross: fns-windows-amd64 fns-darwin-amd64 fns-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/fns-*

fns-linux: fns-linux-386 fns-linux-amd64 fns-linux-mips64 fns-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/fns-linux-*

fns-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/fns
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/fns-linux-* | grep 386

fns-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/fns
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/fns-linux-* | grep amd64

fns-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/fns
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/fns-linux-* | grep mips

fns-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/fns
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/fns-linux-* | grep mipsle

fns-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/fns
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/fns-linux-* | grep mips64

fns-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/fns
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/fns-linux-* | grep mips64le

fns-darwin: fns-darwin-386 fns-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/fns-darwin-*

fns-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/fns
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/fns-darwin-* | grep 386

fns-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/fns
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/fns-darwin-* | grep amd64

fns-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/fns
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/fns-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
