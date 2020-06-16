# kratos frame

先创建角色 注册时会分配普通角色  
id 1 管理员  
id 2 用户

#目录结构说明
注：1. app下的服务属于独立目录 可自行实现其他结构的目录


```
│       ├── api               // protobuf 目录  
│       ├── cmd               // 包含main 程序的启动目录
│       └── configs           // 配置文件
│       └── internal          // 内部目录 包含服务的基础操作
│           └── dao           // db操作目录  
|               └── dao.go mc.go redis.go db.go   // 均为初始化操作 
|               └── mc.cache.go dao.bts.go // 使用命令生成的 缓存代码和缓存回溯代码
|               └── models.bts.go models.mc.go // 命令自动生成缓存代码 和缓存回溯代码 所需要的源文件
|               └── user.db.go  user // 模块的 mysql操作， 一个服务下多个模块 应有多个 .db.go
│           └── di            // 依赖注入 文件 修改时 需要修改 wire.go文件来更改依赖注入配置， wire_gen.go 是命令生成的文件
│           └── err_code      // 错误码目录
│           └── model         // 模块的model
│           └── pkg           // 一些工具包
│           └── server        // 包括 http grpc的路由
│           └── service       // protobuf的服务
│       └── runtime           // 存放日志的目录 也可存放其他不需要传到 git上的文件
│       └── test              // 测试文件目录 单个函数的测试文件 更建议放置在和函数的同一目录
```

# 基础组件

## 错误码
1. pkg 目录下， 编写proto文件使用命令 生成ecode.go
```
# generate ecode
kratos tool protoc --ecode api.proto
```

2. app下服务内部使用错误码需要先去注册错误码， 注册需要使用到的错误码， 参照：
```
# 服务 main方法中需要调用此方法
func NewECode() {
	cms := map[int]string{
		err_code.NoLogin.Code():     "用户未登录",
		err_code.LoginErr.Code():    "登录失败",
		err_code.TokenErr.Code():    "Token认证失败",
		err_code.NoUser.Code():      "用户不存在",
		err_code.PhoneErr.Code():    "手机号错误",
		err_code.PassWordErr.Code(): "密码错误",
	}
	ecode.Register(cms)
}
```
3. 使用：下层出现错误 返回指定错误码给上层或客户端，其他可参考 pkg/ecode Test
```
ecode.Cause(UserNotLogin).Code()
ecode.Cause(UserNotLogin).Message()
```

## 日志
1. 可使用kratos 自带的日志， 日志会输出到默认输出目录， kratos 支持的日志格式更多
2. 使用封装日志 pkg/log 目前支持按天分隔， 输入文件路径，输出至不同的目录

使用方式  其他可参照 pkg/log Test
```
log.Info("hi:%s", "kratos")
Log("./log/test", InfoLevel, log.KVString("name", "xiaoming"))
```

## 限流
1. http 使用自适应限流 根据cpu等资源自动调节， 更多参照 kratos
```
# 挂载自适应限流中间件到 bm engine，使用默认配置
limiter := bm.NewRateLimiter(nil)
e.Use(limiter.Limit())
```

2. Grpc 使用自适应限流 根据cpu等资源自动调节， 更多参照 kratos
```
# 挂载自适应限流拦截器到 warden server，使用默认配置
limiter := ratelimiter.New(nil)
ws.Use(limiter.Limit())
```

## 链路跟踪
app 下服务内 注册链路跟踪, http grpc 拦截器有实现的默认的链路跟踪，redis mc 也有， mysql 因为要替换 gorm 需要再gorm操作层 自行实现链路跟踪
注： 其他需要自行实现 参照kratos 中的 链路跟踪 Test
```
// 注册链路跟踪 http, grpc 默认有链路跟踪的拦截 参考文档：https://book.eddycjy.com/golang/grpc/zipkin.html
env.AppID = "User"
zipkin.Init(&zipkin.Config{
    Endpoint:      "http://127.0.0.1:9411/api/v2/spans", // docker run -d -p 9411:9411 openzipkin/zipkin
    BatchSize:     1000,
    Timeout:       xtime.Duration(2 * time.Second),
    DisableSample: true,
})
```

### 参数验证
kratos自带的  validator.v9
```
使用参考
https://s0godoc0org.icopy.site/gopkg.in/go-playground/validator.v9
```
## 熔断
暂未实现