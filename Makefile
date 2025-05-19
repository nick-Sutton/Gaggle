BINARY_NAME=gaggle

build:
	export GOMODCACHE=./backend/bin/modcache && \
	export GOTMPDIR=./backend/bin && \
	GOARCH=amd64 GOOS=darwin go build -o ./backend/bin/${BINARY_NAME}-darwin ./backend/cmd/Gaggle/main.go
	GOARCH=amd64 GOOS=linux go build -o ./backend/bin/${BINARY_NAME}-linux ./backend/cmd/Gaggle/main.go
	GOARCH=amd64 GOOS=windows go build -o ./backend/bin/${BINARY_NAME}-windows ./backend/cmd/Gaggle/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean -cache
	rm -f ./backend/bin/${BINARY_NAME}-*
