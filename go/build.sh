cd ../
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
protoc --proto_path=. --go_out=go/internal/runtime/generated --go_opt=paths=source_relative src/ray/protobuf/common.proto
