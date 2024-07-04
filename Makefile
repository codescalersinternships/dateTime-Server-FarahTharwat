GIN_DIR=cmd/gin_server
STD_DIR=cmd/std_server
GIN_BIN=ginserver
STD_BIN=stdserver

build-gin:
	GOARCH=amd64 GOOS=darwin go build -o ${GIN_BIN}-darwin ${GIN_DIR}/main.go 
	GOARCH=amd64 GOOS=linux go build -o ${GIN_BIN}-linux ${GIN_DIR}/main.go
	GOARCH=amd64 GOOS=windows go build -o ${GIN_BIN}-windows ${GIN_DIR}/main.go

build-std:
	GOARCH=amd64 GOOS=darwin go build -o ${STD_BIN}-darwin main.go 
	GOARCH=amd64 GOOS=linux go build -o ${STD_BIN}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${STD_BIN}-windows main.go
 # go releaser saves us from building for different oses and architectures 

build: build-gin build-std

fmt: 
	go fmt ./...

lint: 
	golangci-lint run 
test:
	go test -v ./...

