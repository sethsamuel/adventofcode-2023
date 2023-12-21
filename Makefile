BINARY_NAME=adventofcode-2023

build:
	go build -o ${BINARY_NAME}.exe main.go

run: build
	./${BINARY_NAME}.exe

clean:
	go clean
	rm ${BINARY_NAME}

day = 1
test:
	go test ./day$(day)/...

debug:
	go test -v ./day$(day)/...