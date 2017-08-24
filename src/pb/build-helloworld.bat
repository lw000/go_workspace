cd ../..
set GOPATH=%cd%
cd src/pb
protoc.exe --go_out=plugins=grpc:./helloworld ./helloworld.proto                                        