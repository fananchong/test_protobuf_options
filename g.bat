go list -m -f "{{.Dir}}" github.com/fananchong/test_protobuf_options > temp
set /P DEP1=<temp
go list -m -f "{{.Dir}}" github.com/golang/protobuf > temp
set /P DEP2=<temp
echo %DEP1%
echo %DEP2%

protoc -I. -I%DEP1% -I%DEP2% --go_out=. test.proto