build:
	go build -o dist/broker cmd/broker/main.go
	go build -o dist/sub cmd/sub/main.go
	go build -o dist/pub cmd/pub/main.go

clean:
	go clean
	rm -rf dist
