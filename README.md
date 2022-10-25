# aws-glados-checkin
GLADOS Automatic Check-in -- AWS lambda Version

* Author: [xiabee](https://github.com/xiabee)
* Update time: 2022.10.25
* Version: 1.0



## 项目功能

* 基于云函数签到
* 基于 `Git Action` 的签到：https://github.com/xiabee/glados-checkin

* 利用 [AWS lambda](https://ap-southeast-1.console.aws.amazon.com/lambda/home?region=ap-southeast-1) 进行云函数编程，实现[Glados](https://glados.rocks/console) 自动签到，并将签到结果推送至企业微信

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7hz51nsbbj31pq0rmty1.jpg)



## 编译模式

`lambda go` 仅支持 `x86` 架构，因此在编译二进制文件时需要添加 `GOOS=linux GOARCH=amd64` 参数

```bash
# Remember to build your handler executable for Linux!
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```



## 使用方式

待更新
