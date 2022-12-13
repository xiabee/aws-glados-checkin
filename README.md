# aws-glados-checkin
GLADOS Automatic Check-in -- AWS lambda Version

* Author: [xiabee](https://github.com/xiabee)
* Update time: 2022.12.13
* Version: 1.0



## 项目功能

* 基于 AWS 云函数签到
* 利用 [AWS lambda](https://ap-southeast-1.console.aws.amazon.com/lambda/home?region=ap-southeast-1) 进行云函数编程，实现[Glados](https://glados.rocks/console) 自动签到，并将签到结果推送至企业微信

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7hz51nsbbj31pq0rmty1.jpg)

* 为什么放弃基于 `Git Action` 的签到：https://github.com/xiabee/glados-checkin
  * `GitHub` 曾经对利用 `Action` 签到的仓库进行封禁，当今使用 `Action` 进行签到等功能不是很安全
  * AWS Lambda 对于签到这个功能来说，基本就是免费（因为每个月免费额度很多，仅靠签到是完全用不完的）

## 编译模式

`lambda go` 仅支持 `x86` 架构，因此在编译二进制文件时需要添加 `GOOS=linux GOARCH=amd64` 参数

```bash
# Remember to build your handler executable for Linux!
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```



## 使用方式

* 直接编译上传至 AWS Lambda，程序入口点设置为`main`

* 设置一下 Glados 的 `COOKIE`，COOKIE 获取过程参考`附录`
* 如果需要连接企业微信机器人，可以设置一下`WECHAT_KEY`，具体的机器人设置参考本篇：https://blog.xiabee.cn/posts/wechat-bot/

* AWS Lambda 详细使用过程参考本篇：https://blog.xiabee.cn/posts/aws-lambda/



## 附录

### Cookie 获取方式

* 登录[控制台](https://glados.rocks/console)之后，`F12`打开浏览器的开发者工具
* 找到 `Network` / `网络`，选中 `console` 那个包
* 在请求头中找到 `cookie` 字段

<img src="https://tva1.sinaimg.cn/large/0084b03xgy1h7f4k2lqajj31vo0zk1iq.jpg" alt="image.png" style="zoom:67%;" />
