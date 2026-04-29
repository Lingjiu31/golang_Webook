.PHONY: docker
docker:
#	删掉之前的,防止冲突
	@rm -f webook
#	编译 Linux 版程序
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=k8s -o webook ./cmd/server
#	删掉之前的,防止冲突
	@docker rmi -f lingjiu31/webook:v0.0.1
#	构建 docker 镜像
	@docker build -t lingjiu31/webook:v0.0.1 .
#	触发 k8s 滚动更新
	@kubectl rollout restart deployment webook
