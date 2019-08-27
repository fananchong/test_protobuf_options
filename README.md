# test_protobuf_options
使用 protobuf 的自定义选项功能


## 制作自定义选项说明

本例子演示如何使用 protobuf 的自定义选项功能，给 rpc 方法加上 broadcast 选项

包含 2 部分内容：

- 定义 broadcast 选项
  - 请参考 broadcast.proto 方式定义

- 基于 proto_gen_go 自定义插件，根据 broadcast 选项来生成自定义代码
  - protoc-gen-test 目录代码实现

## 如何使用自定义选项说明

protobuf 自定义选项机制，使得例子不能在同一个项目内，因此例子参考：

https://github.com/fananchong/test_protobuf_options_example


该例子主要定义了 test.proto 文件：

```protobuf
syntax = "proto3";

import "broadcast.proto";

package proto;

service Say {
    rpc Hello(Request) returns (mypack.NoReply) { option (mypack.Broadcast) = true; }
    rpc Ping(Request) returns (mypack.NoReply) { }
}

message Request {
    string name = 1;
}
```

并通过本项目自定义插件 proto_gen_test ，实现代码输出 test.pb.go ：

```go
// 篇幅问题等，注释掉同 proto_gen_go 生成代码 略

// For example

type SayService interface {
	BroadcastHello(ctx context.Context) error
	Ping(ctx context.Context) error
}
```
