build:
	go build -o bin/resv

run: build
	./bin/resv

test:
	go test -v ./...
