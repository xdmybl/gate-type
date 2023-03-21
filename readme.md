# what
gate-type 是 engine-gate 的一个工具, 用于生成 engine-gate 的 gate resource 的 CRD 和 golang type


# why
这个项目专门保存这些资源, 可以不用影响到 engine-gate 项目


# how
## how to implement
使用 protoc 工具和配套的生态, 根据 .proto 文件定义好标准类型, 然后根据 .proto 文件生成对应的 CRD 定义文件 和 golang type.
```shell
protoc --go_out="paths=source_relative:tmp" -I . --proto_path=/home/xiaoy/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7 proto/core/v1/base.proto
```

## how to install env to dev
执行 make install_tools

## how to use
在改变了 proto 目录下的 proto 文件后, 执行 make run , tmp 目录将出现生成的 golang 代码, 然后执行 make copy , 将 tmp 代码拷贝 api 目录中

# now 
## 各工具版本及下载链接
    protoc v3.18.1  https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protoc-3.18.1-linux-x86_64.zip
    protoc-gen-go v1.28.1  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1

## 问题