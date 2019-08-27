set CUR_DIR=%~dp0
set GOPROXY=https://goproxy.io
set GOBIN=%CUR_DIR%bin
go install ./...