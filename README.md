## 微信小程序商城-服务端

开源的小程序：https://github.com/EastWorld/wechat-app-mall

开源CMS：https://www.talelin.com/

## Token双令牌机制

#### access_token 和 refresh_token

jwt机制颇为灵活，如Github选择了单令牌机制，lin-cms为了增强用户体验、提高接口安全性， lin-cms选择了双令牌机制。在双令牌机制中，
access_token和refresh_token是一对相互帮助的好搭档，用户登陆成功后，服务器 会颁发 access_token 和 refresh_token，前端在
得到这两个token之后必须谨慎存储，它们的作用如下：

- access_token：用户访问接口，资源的凭证。access_token十分重要，它是服务器对前端有力控制的唯一途径，其生存周期较短，一般在2个
小时左右，更有甚者，其生命周期只有15分钟。

- refresh_token：用户重新获得access_token的凭证。refresh_token的生命周期较长 ，一般为30天左右，但refresh_token不能被用
来用户身份鉴权和获取资源，它只能被用来重新获取access_token。当前端发现access_token过期时，会自动通过 refresh_token重新获取
access_token

## Road Map

1.前后端分离-API接口（小程序、CMS）

2.调度任务

