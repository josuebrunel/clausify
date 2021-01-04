test:
	go test -v -count=1 -cover -covermode=count -coverprofile=coverage.out ./...
lint:
	golint ./...
clean:
	go clean
debug:
	dlv test github.com/josuebrunel/clausify
all: test lint
