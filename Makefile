BUF_VERSION:=v1.17.0

generate: go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate

clean:
	rm pb/*
	rm bin/*

server:
	@go build -o bin/server main.go
	@./bin/server