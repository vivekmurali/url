build:
	go build -o url.out .

test:
	go test -v ./...

run:
	./url.out

br: build run
