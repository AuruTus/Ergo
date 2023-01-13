
.PHONY: test
test:
	go test ./...


# used for test specefic function
# `func` is the go test function name like `TestNewConfiguredLogger` 
.PHONY: test-func
test-func:
	go test -v -run ${func} ./...
