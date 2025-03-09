CMD_DIR := ./cmd

get:
	go get $(CMD_DIR)
run:
	make swag-init && go run $(CMD_DIR)/main.go
build:
	make swag-init && make clean && go build -C $(CMD_DIR) -o ../bin/go-redis-sample
start:
	make build && cp .env ./bin/.env && ./bin/go-redis-sample
clean:
	rm -rf ./bin/
swag-init:
	go install github.com/swaggo/swag/cmd/swag@latest && swag init -g $(CMD_DIR)/main.go -o docs
