cd ../../
set GOPATH=%cd%
cd src/SocketServer0001
set GOARCH=amd64
set GOOS=linux
go clean
go build