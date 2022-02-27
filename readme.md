# goal-wechat
这是一个 goal 的阿里云 sdk，基于 [denverdino/aliyungo](https://github.com/denverdino/aliyun) 封装，支持多应用

## 安装 - install
```shell
$ go get github.com/qbhy/goal-aliyun
```

## 配置 - configuration
```go
// config/aliyun.go
```

## 使用 - usage
注册服务
```go
// main.go
import (
    "github.com/goal-web/application"
    aliyun "github.com/qbhy/goal-aliyun"
)

func main()  {
    var app = application.Singleton()
    
    app.RegisterServices(
        // other service
        aliyun.ServiceProvider{},
    )
	// ...
}
```
使用
```go
// app/controllers/order/pay.go

import (
	"fmt"
    "github.com/goal-web/contracts"
    aliyun "github.com/qbhy/goal-aliyun"
)

// ... 

// ...
```

## 相关链接  
[goal-web](https://github.com/goal-web/goal)  
[qbhy/goal-aliyun](https://github.com/qbhy/goal-aliyun)  
qbhy0715@qq.com
