gen-cal:
	protoc api/proto/movie.proto --go-grpc_out=.
	protoc api/proto/movie.proto --go_out=.
	
run-server:
	go run cmd/main.go
run-client:
	go run client/main.go	