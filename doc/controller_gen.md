# 目前已在项目中废弃其使用, 换成 skv2

# 概述
controller-gen 是 controller-tools 里面的工具, 本项目中主要用于根据 skv2 生成的 .go 文件生成 kubernetes crd yaml 文件.

# 使用方法
```shell
# crd 是子命令    paths 是传入的
controller-gen crd paths=./... output:crd:dir=crds
```

