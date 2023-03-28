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


.PHONY: gen
gen_type:
	@echo "start generate golang type using skv2"
	@echo "api/ : 保存 common 和 core 的对象"
	@echo "pkg/api/ : 保存 gate 的对象"
	@go run main.go


# 在有删除对象操作时一定要执行这个, 因为有些 pb.go 文件会被删除掉.
.PHONY: regenerate
regenerate:
	@echo "正在删除之前的文件...... "
	@rm -rf pkg/api/*
	@rm -f crds/*
	@echo "正在重新生成 golang type 和 crd 文件"
	@go run main.go
