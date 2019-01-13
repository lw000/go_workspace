cd ../../../../
set GOPATH=%cd%
cd src/Gate/cmd/Gate
set GOARCH=amd64
set GOOS=linux
go clean
go build