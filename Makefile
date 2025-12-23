run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test ./...

tidy:
	go mod tidy

fmt:
	go fmt ./...

clean:
	rm -rf bin

lint:
	go vet ./...
