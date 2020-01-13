run_test:
	go run main.go -dir=example -name=test

build_linux_x64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o builds/cbuild_linux_x64 .
