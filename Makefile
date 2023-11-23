gen-cal:
	protoc proto/movie.proto --go-grpc_out=.
	protoc proto/movie.proto --go_out=.
	
run-server:
	go run server/main.go
run-client:
	go run client/main.go	