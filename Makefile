
ROOTPATH:=${realpath .}

ERGO_MAIN:=${ROOTPATH}/cmd/ergo/
ERGO_EXE:=${ROOTPATH}/bin/ergo/go-ergo

ERGO_SRC:=${shell find ./cmd/ergo -type f -name '*'} \
${shell find ./pkg -type f -name '*'} \
${shell find ./internal -type f -name '*'}


.PHONY: build
build: ${ERGO_SRC}
	go build -o ${ERGO_EXE} ${ERGO_MAIN}

${ERGO_EXE}: ${ERGO_SRC}
	go build -o $@ ${ERGO_MAIN}

.PHONY: run
run: ${ERGO_EXE}
	@$<


.PHONY: test
test:
	go test ./...


# used for test specefic function
# `func` is the go test function name like `TestNewConfiguredLogger` 
.PHONY: test-func
test-func:
	go test -v -run ${regex} ./...


.PHONY: test-func-bench
test-func-bench:
	go test -v -run XXX -bench ${regex} ./...
