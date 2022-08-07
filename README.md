## 微信小程序商城-服务端

本项目是 [wechat-mall](https://github.com/ZuoFuhong/wechat-mall-miniapp) 微信小程序商城 配套的`服务端`

### Development

```sh
# clone the project
git clone git@github.com:ZuoFuhong/wechat-mall-backend.git

# update dependency
go mod tidy

# build the project
make

# init the database
/doc/init_wechat_mall.sql
```

### Deployment

1.在项目的[Release](https://github.com/ZuoFuhong/wechat-mall-backend/releases)页面，下载最新的包 wechat-mall-backend-${last-version}.tar.gz

2.初始化数据库，下载项目 doc 目录下的`init_wechat_mall.sql`脚本（默认连接的数据库是`wechat_mall`）

3.运行服务端前，需要修改配置文件，配置基本参数（redis、mysql、阿里云OSS、小程序配置等）

4.本项目是一个web项目，默认运行在`8080`端口，启动命令：

```
$ ./wechat-mall-backend    # 根据环境变量 RUN_MODE，读取配置文件，默认读取 dev 环境 
```


### RoadMap

- 1.店铺信息设置
- 2.会员体系
- 3.访客记录
- 4.物流模块，运费模板
- 5.订单列表（发货与退款）
- 6.售罄的商品
- 7.小程序分享，绘制canvas海报
- 8.后端调度任务（退款调度、待付款调度）

### 最佳实践

- [腾讯云微搭低代码开发平台](https://cloud.tencent.com/product/weda) 推荐的最佳应用案例 [【云开发应用】麦兜小店-电商小程序](https://github.com/WeDaHub/incubator-mcdull-mall).

### Backer and Sponsor
> jetbrains

<a href="https://www.jetbrains.com/?from=ZuoFuhong/bulb" target="_blank">
<img src="https://github.com/ZuoFuhong/bulb/blob/master/doc/jetbrains-gold-reseller.svg" width="100px" height="100px">
</a>

### License

The project is licensed under the Apache 2 license.
