go list -m -f "{{.Dir}}" github.com/golang/protobuf > temp
set /P DEP=<temp
echo %DEP%
protoc -I. -I%DEP% --go_out=. broadcast.proto