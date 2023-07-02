clean:
	rm pb/*
	rm bin/*

gen:
	protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb 

server:
	@go build -o bin/server cmd/main.go
	@./bin/server