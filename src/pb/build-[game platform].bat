cd ../..
set GOPATH=%cd%
cd src/pb
protoc --go_out=./game ./game.proto
protoc --go_out=./platform ./platform.proto