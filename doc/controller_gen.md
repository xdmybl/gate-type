# 概述
controller-gen 是 controller-tools 里面的工具, 本项目中主要用于根据 skv2 生成的 .go 文件生成 kubernetes crd yaml 文件.

# 使用方法
```shell
# crd 是子命令    paths 是传入的
controller-gen crd paths=./... output:crd:dir=crds
```

# TODO
目前执行 controller-gen 命令后还有报错, 尚未解决, 不过生成的 crd yaml 文件没什么问题. 后续解决报错