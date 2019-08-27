go list -m -f "{{.Dir}}" github.com/golang/protobuf > temp
set /P DEP=<temp
echo %DEP%
protoc -I. -I%DEP% --go_out=. broadcast.proto
copy /y github.com\fananchong\test_protobuf_options\broadcast.pb.go .
rd /s /q github.com
