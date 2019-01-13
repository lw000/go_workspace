cd ../../
set GOPATH=%cd%
cd src/SocketClientTest
set GOARCH=amd64
set GOOS=linux
go clean
go build