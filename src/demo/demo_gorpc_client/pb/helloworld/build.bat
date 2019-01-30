cd ../../../..
set GOPATH=%cd%
cd src/demo_gorpc_client/pb/helloworld
protoc.exe --go_out=plugins=grpc:. ./*.proto                                        