cd Services/protos

protoc --go_out=../../ Models.proto
protoc --micro_out=../../ --go_out=../../ TestService.proto
protoc-go-inject-tag -input=../../Services/Models.pb.go
protoc-go-inject-tag -input=../../Services/TestService.pb.go

cd ..  & cd ..