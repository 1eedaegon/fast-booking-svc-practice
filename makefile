build:
	go build -o bin/resv

run: build
	./bin/resv

test:
	go test -v ./...

run-mongo:
	docker run --name mongo -d -p 27017:27017 mongo
