set GOPATH=%cd%
set GOBIN=%cd%\bin
cd src
go install main.go
cd ..
exit