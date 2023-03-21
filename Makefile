.PHONY: install_tools
install_tools:
	# 存在 protoc 命令则不下载
	@if [ ! `command -v protoc` ]; then  wget https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protoc-3.18.1-linux-x86_64.zip && unzip protoc-3.18.1-linux-x86_64.zip -d protoc3  && sudo mv protoc3/bin/* /usr/local/bin/ && mkdir -p ${HOME}/include/ && mv protoc3/include/* ${HOME}/include/ && rm -rf protoc-3.18.1-linux-x86_64.zip protoc3; else echo "already exist"; fi
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
	@protoc --version

.PHONY: tidy
tidy:
	go mod tidy


PROTO_FILES := $(shell find proto -name "*.proto" -type f)

.PHONY: run
run:
	# 这个指令报错的话可能是因为 protoc-gen-validate 版本不合适. 可参考 上面 install_tools
	# 这个 for 很容易错, 主要是 ; 要加上, 不然会报错
	@for file in $(PROTO_FILES); do \
		echo "根据 proto 文件  $$file, 生成 golang 文件"; \
		protoc --go_out="paths=source_relative:tmp" -I . -I ${HOME}/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7 $$file ; \
    done
	@echo "完成 golang type 的生成"
	@echo "正在拷贝到 api 目录中"
	@cp -r tmp/proto/* api



	#go run api/generate.go
	# the api/generate.go command is separated out to enable us to run go generate on the generated files (used for mockgen)
## this re-gens test protos
#	go test ./codegen
#	go generate -v ./...