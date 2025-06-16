OUTPUT_DIR := dist

.DEFAULT_GOAL := build

ARCH_MAP_arm64 := arm64

build: dot-darwin-arm64

codegen: lib/common/cache/Colfer.go

dot-%: codegen
	@GOOS=$(word 1,$(subst -, ,$*)) \
	GOARCH=$(ARCH_MAP_$(word 2,$(subst -, ,$*))) \
	go build -o $(OUTPUT_DIR)/dot-$*

lib/common/cache/Colfer.go: lib/common/cache/cache.colf
	colf -b lib/common Go lib/common/cache/cache.colf

.PHONY: build
