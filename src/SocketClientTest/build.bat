cd ../../
set GOPATH=%cd%
cd src/SocketClientTest
set GOARCH=amd64
set GOOS=windows
go clean
go build