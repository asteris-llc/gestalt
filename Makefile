TEST?=./...
NAME = $(shell awk -F\" '/^const Name/ { print $$2 }' main.go)
VERSION = $(shell awk -F\" '/^const Version/ { print $$2 }' main.go)

all: deps test

gestalt: deps
	go build -o gestalt .

deps:
	go get -v -t ./...
	go get -v github.com/raphael/goa/goagen

test:
	go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4
	go vet $(TEST)

# web

web/app: web/design/*.go
	cd web && goagen app -d github.com/asteris-llc/gestalt/web/design

web/swagger: web/design/*.go
	cd web && goagen swagger -d github.com/asteris-llc/gestalt/web/design

web/impl: web/design/*.go
	@test -d web/impl || mkdir web/impl
	cd web/impl && goagen main -d github.com/asteris-llc/gestalt/web/design

web/web.go: web/impl web.go.sed
	echo '// !!! automatically generated !!!\n// Use "make web/web.go" instead of editing this file.\n' | cat - web/impl/main.go > web/web.go
	sed -E -i '' -f web.go.sed web/web.go
	gofmt -w web/web.go

docs/cli: cmd/**/*.go *.go
	go run *.go __markdown

# packaging

xcompile: deps test
	@rm -rf build/
	@mkdir -p build
	gox \
		-os="darwin" \
		-os="dragonfly" \
		-os="freebsd" \
		-os="linux" \
		-os="openbsd" \
		-os="solaris" \
		-os="windows" \
		-output="build/{{.Dir}}_$(VERSION)_{{.OS}}_{{.Arch}}/$(NAME)"

package: xcompile
	$(eval FILES := $(shell ls build))
	@mkdir -p build/tgz
	for f in $(FILES); do \
		(cd $(shell pwd)/build && tar -zcvf tgz/$$f.tar.gz $$f); \
		echo $$f; \
	done

.PHONY: all deps test xcompile package
