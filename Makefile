# https://habr.com/ru/post/461467/
# https://tutorialedge.net/golang/makefiles-for-go-developers/
# https://earthly.dev/blog/golang-makefile/
BINARY_NAME=csv2md
ARCH=amd64

LINUX_PATH="bin/linux_${ARCH}"
WINDOWS_PATH="bin/windows_${ARCH}"
DARWIN_PATH="bin/darwin_${ARCH}"

LINUX_FILE="${LINUX_PATH}/${BINARY_NAME}"
WINDOWS_FILE="${WINDOWS_PATH}/${BINARY_NAME}.exe"
DARWIN_FILE="${DARWIN_PATH}/${BINARY_NAME}"

## clean: Removes all compiled binaries
clean:
	@go clean
	@rm -rf bin/

## linux: Builds new binaries for linux (x64)
linux:
	@rm -rf ${LINUX_PATH}
	@GOARCH=${ARCH} GOOS=linux go build -o ${LINUX_FILE} . && echo "Compiled: ${LINUX_FILE}"

## win: Builds new binaries for windows (x64)
win:
	@rm -rf ${WINDOWS_PATH}
	@GOARCH=${ARCH} GOOS=windows go build -o ${WINDOWS_FILE} . && echo "Compiled: ${WINDOWS_FILE}"

## darwin: Builds new binaries for darwin (x64)
darwin:
	@rm -rf ${DARWIN_PATH}
	@GOARCH=${ARCH} GOOS=darwin go build -o ${DARWIN_FILE} . && echo "Compiled: ${DARWIN_FILE}"

## build: Builds new binaries for linux, windows and darwin (x64)
all: clean linux win darwin

## compile: This message
help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
