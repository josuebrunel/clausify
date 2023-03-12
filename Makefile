test:
	go test -v -failfast -count=1 -cover -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func coverage.out
test.integration:
	docker compose -f build/docker-compose.yml up --build -d
	cd tests/; go mod tidy; go test -v -count=1 -cover -covermode=count -coverprofile=coverage.out ./... ; go tool cover -func coverage.out ; cd -
lint:
	golint ./...
clean:
	go clean
debug:
	dlv test github.com/josuebrunel/clausify

all: clean test test.integration
