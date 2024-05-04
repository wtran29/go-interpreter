make build:
	go build -o funckey . && ./funckey

test_compiler:
	go test ./src/compiler

test_vm:
	go test ./src/vm