get:
	go get ./cmd/
run:
	go run ./cmd/main.go
build:
	make clean && go build -C cmd -o ../bin/go-redis-sample
clean:
	rm -rf bin/
