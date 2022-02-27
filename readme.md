# goal-wechat
这是一个 goal 的支付宝 sdk，基于 [smartwalle/alipay](https://github.com/smartwalle/alipay) 封装，支持多应用

## 安装 - install
```shell
$ go get github.com/qbhy/goal-alipay
```

## 配置 - configuration
```go
// config/alipay.go
package config

import (
	"github.com/goal-web/contracts"
	alipay "github.com/qbhy/goal-alipay"
)

func init() {
	configs["alipay"] = func(env contracts.Env) interface{} {
		return &alipay.Config{
			Default: env.StringOption("alipay.default", "default"),
			Apps: map[string]*alipay.AppConfig{
				"default": {
					AppId:        env.GetString("alipay.appid"),
					PrivateKey:   env.GetString("alipay.private_key"),
					IsProduction: !env.GetBool("alipay.debug"),
					//OptionFunctions: []alipay.OptionFunc{
					//	func(c *alipay2.Client) {
					//		// do something
					//	},
					//},
				},
			},
		}
	}
}

```

## 使用 - usage
注册服务
```go
// main.go
import (
    "github.com/goal-web/application"
alipay "github.com/qbhy/goal-alipay"
)

func main()  {
    var app = application.Singleton()
    
    app.RegisterServices(
        // other service
        alipay.ServiceProvider{},
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
    alipay "github.com/qbhy/goal-alipay"
    alipayWalle "github.com/smartwalle/alipay/v3"
)

// ... 

// 通过注入支付宝工厂，你可以通过工厂获取定义好的应用
func AppTrade(alipay alipay.Factory, request contracts.HttpRequest) {
    var client = alipay.Client("default")
}

// 你也可以直接注入 支付宝kehuduan  的实例，这种情况下会获取默认实例
func InjectInstanceExample(client *alipayWalle.Client) {
	// do something ...
}

// ...
```

## 相关链接  
[goal-web](https://github.com/goal-web/goal)  
[qbhy/goal-alipay](https://github.com/qbhy/goal-alipay)  
[smartwalle/alipay](https://github.com/smartwalle/alipay)  
qbhy0715@qq.com
