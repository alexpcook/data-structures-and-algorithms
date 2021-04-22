BINARY_NAME=main.out
TEST_PACKAGES="./..."

all: test

build:
	go build -o ${BINARY_NAME}

test:
	go test -v -bench . -benchmem ${TEST_PACKAGES}

run:
	go build -o ${BINARY_NAME}
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
