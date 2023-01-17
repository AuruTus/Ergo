

# TODO: add init config for project
# .PHONY: init-config
# init-config:
# # if the project is new cloned from remote repo, the tstr_entry.go will be missing
# ifeq (,$(wildcard ./tstr_entry.go))
# 	./scripts/init_dojo.sh
# endif


# build: init-config
build:
	go build -o ./bin/ergo ./cmd/ergo


.PHONY: run
run: build
	./bin/ergo


.PHONY: test
test:
	go test ./...


# used for test specefic function
# `func` is the go test function name like `TestNewConfiguredLogger` 
.PHONY: test-func
test-func:
	go test -v -run ${func} ./...
