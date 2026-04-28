.PHONY: docker
docker:
#	删掉之前的,防止冲突
	@rm webook || ture
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webook .
#	删掉之前的,防止冲突
	@docker rmi -f lingjiu31/webook:v0.0.1
	@docker build -t lingjiu31/webook:v0.0.1 .
