test:
	go test -v -cover ./...

server:
	go run .


.PHONY: test server