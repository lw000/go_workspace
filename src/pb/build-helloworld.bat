cd ../..
set GOPATH=%cd%
cd src/pb
protoc --go_out=plugins=grpc:./helloworld ./helloworld.proto