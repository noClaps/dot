OUTPUT_DIR := dist

.DEFAULT_GOAL := build

ARCH_MAP_arm64 := arm64

build: doot-darwin-arm64

codegen: lib/common/cache/Colfer.go

doot-%: codegen
	@GOOS=$(word 1,$(subst -, ,$*)) \
	GOARCH=$(ARCH_MAP_$(word 2,$(subst -, ,$*))) \
	go build -o $(OUTPUT_DIR)/doot-$*

lib/common/cache/Colfer.go: lib/common/cache/cache.colf
	bin/colf -b lib/common Go lib/common/cache/cache.colf

.PHONY: build
