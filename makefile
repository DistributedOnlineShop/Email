test:
	go test -v -cover ./...

server:
	go run .

protob:
	rm -f pb/email/*.go
	rm -f pb/verification/*.go
	rm -f pb/session/*.go
	rm -f pb/users/*.go

	protoc --proto_path=proto --go_out=pb/email --go_opt=paths=source_relative \
	--go-grpc_out=pb/email --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/email --grpc-gateway_opt paths=source_relative \
    proto/service_email.proto

	protoc --proto_path=proto --go_out=pb/verification --go_opt=paths=source_relative \
	--go-grpc_out=pb/verification --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/verification --grpc-gateway_opt paths=source_relative \
    proto/service_verification.proto

	protoc --proto_path=proto --go_out=pb/session --go_opt=paths=source_relative \
	--go-grpc_out=pb/session --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/session --grpc-gateway_opt paths=source_relative \
    proto/service_session.proto

	protoc --proto_path=proto --go_out=pb/users --go_opt=paths=source_relative \
	--go-grpc_out=pb/users --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/users --grpc-gateway_opt paths=source_relative \
    proto/service_users.proto

evans: 
	evans --host localhost --port 8193 -r repl

.PHONY: test server protob evans
