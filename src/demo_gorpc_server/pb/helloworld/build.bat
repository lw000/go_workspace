cd ../../../..
set GOPATH=%cd%
cd src/demo_gorpc_server/pb/helloworld
protoc.exe --go_out=plugins=grpc:. ./*.proto                                        