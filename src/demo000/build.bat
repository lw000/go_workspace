cd ../../
set GOPATH=%cd%
cd src/demo000
set GOARCH=amd64
set GOOS=windows
go clean
go build