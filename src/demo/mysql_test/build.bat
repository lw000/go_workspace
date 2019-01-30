set GOPATH=%cd%\..\..
set GOBIN=%cd%\bin
cd mysql_test
go install main.go
cd ..
exit                                                  