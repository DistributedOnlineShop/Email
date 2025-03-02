test:
	go test -v -cover ./...

server:
	go run .

sProtob:
	rm -f pb/server/*.go
	rm -f pb/client/*.go

	protoc --proto_path=proto --go_out=pb/grpc_server --go_opt=paths=source_relative \
	--go-grpc_out=pb/grpc_server --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/grpc_server --grpc-gateway_opt paths=source_relative \
    proto/server.proto

	protoc --proto_path=proto --go_out=pb/grpc_client --go_opt=paths=source_relative \
	--go-grpc_out=pb/grpc_client --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/grpc_client --grpc-gateway_opt paths=source_relative \
    proto/client.proto


evans: 
	evans --host localhost --port 8193 -r repl

.PHONY: test server sProtob evans