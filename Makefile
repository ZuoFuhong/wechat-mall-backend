NAME = wechat-mall-mall
VERSION = 1.0.0
GOOS = linux
GOARCH = amd64

.PHONY: build

build:
	@echo "build ${NAME}-${VERSION}"
	@export GOOS=${GOOS}; \
	export GOARCH=${GOARCH}; \
	go build -o ${NAME} main.go
	@mkdir ${NAME}-${VERSION}
	@mkdir ${NAME}-${VERSION}/conf
	@cp -r conf/ ${NAME}-${VERSION}/conf
	@mv ${NAME} ${NAME}-${VERSION}/
	@tar -zcvf ${NAME}-${VERSION}.tar.gz ${NAME}-${VERSION}; rm -rf ${NAME}-${VERSION}
	@echo "build Done."
