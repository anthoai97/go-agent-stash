BUF_VERSION:=v1.17.0
SWAGGER_UI_VERSION:=v4.15.5

generate: generate/proto

generate/proto:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate

clean:
	rm pb/*
	rm bin/*

gen:
	protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb 

server:
	@go build -o bin/server main.go
	@./bin/server