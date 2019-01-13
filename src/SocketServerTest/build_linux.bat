cd ../../
set GOPATH=%cd%
cd src/SocketServerTest
set GOARCH=amd64
set GOOS=linux
go clean
go build