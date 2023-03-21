#### protoc 执行报错  xxx: Import "protobuf/core/v1/types.proto" was not found or had errors.
    根据报错信息, 某个文件某一行报错, 进到文件去排查, import 报错说明找不到 本地库 或者 某官方库 或者 第三方库.
    通过 -I $path 可以指定库的导入路径, 需要能够通过 $path 和 import 的路径找到对应的文件才行. 注意: 一定要在前面加上 -I . 把本地目录加载上去 

#### go mod tidy 报错中止
    github.com/xdmybl/gate-type/tmp/proto/common/v1 imports
        github.com/xdmybl/gate-type/api/core/v1: no matching versions for query "latest"

    找到相应的行, 先注释 import 的那一行, 再 go mod tidy, 即可. 

