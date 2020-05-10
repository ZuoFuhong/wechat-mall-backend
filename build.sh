#!/usr/bin/env bash

set -ex
echo "build wechat-mall-backend ..."
if [[ -d "target/" ]]
then
rm -rf target/
else
mkdir target
fi
# linux
GOOS=linux GOARCH=amd64 go build -o ./target/wechat-mall-backend main.go
# windows
#GOOS=windows GOARCH=amd64 go build -o ./target/wechat-mall-backend main.go
# macOS
#go build -o ./target/wechat-mall-backend main.go
cp -rf conf ./target/
tar -zcvf ./target/wechat-mall-backend-1.0.0.tar.gz target/
echo "build Done."