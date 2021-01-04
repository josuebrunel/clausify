test:
	go test -v -cover ./...
lint:
	golint ./...
clean:
	go clean
debug:
	dlv test github.com/josuebrunel/clausify
all: test lint
