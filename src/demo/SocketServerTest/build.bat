cd ../../
set GOPATH=%cd%
cd src/SocketServerTest
set GOARCH=amd64
set GOOS=windows
go clean
go build