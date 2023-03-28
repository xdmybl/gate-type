# 注意点
写一个 protocol buffers 文件, 一定要注意 package 和 option go_package 选项

### package  
指示这个文件属于哪个包

### option go_package
`option go_package` 是一个 Protocol Buffers 的选项，用于指定生成的 Go 代码的包名。这个选项告诉编译器将生成的 Go 代码放在哪个 Go 包里面。

例如，如果你有一个名为 `example.proto` 的 Protocol Buffers 文件，并且你想要生成的 Go 代码放在 `github.com/example/protobuf` 这个包中，你可以在 `example.proto` 文件中添加以下行：

```
syntax = "proto3";

package example;

option go_package = "github.com/example/protobuf";
```

这将告诉编译器将生成的 Go 代码放在 `github.com/example/protobuf` 这个包中。

