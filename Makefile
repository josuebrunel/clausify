test:
	go test -v -count=1 -cover -covermode=count -coverprofile=coverage.out ./...
test.integration:
	docker-compose -f build/docker-compose.yml up --build -d
	cd tests/; go test -v -count=1 -cover -covermode=count -coverprofile=coverage.out ./...; cd -
lint:
	golint ./...
clean:
	go clean
debug:
	dlv test github.com/josuebrunel/clausify
all: test lint
