set GOPATH=%cd%\..\..
set GOBIN=%cd%\bin
cd melody_test
go install main.go
cd ..
exit                                                  