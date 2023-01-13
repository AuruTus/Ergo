
.PHONY: test
test:
	go test ./...

.PHONY: test-func
test-func:
	go test -v -run ${func} ./...
