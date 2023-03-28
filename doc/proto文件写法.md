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

### common core gate 目录使用方法
core 里面存放只被引用的类型, 从不引用项目中其它文件
common 中存放只引用 core 的类型
gate 存放所有的 spec 对象
把必须抽取的放到一个文件内, 能合并到一个文件的尽量放一个文件里, 因为同级目录下的 proto 不能直接引用

### spec proto 文件的定义示例
1.     
```shell
# 引用 types
import "proto/core/v1/types.proto";
# 引用 commonInfo , 这是没个 spec 都有的
core.v1.CommonInfo commonInfo = 1;
``` 
