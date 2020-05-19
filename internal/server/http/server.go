package http

import (
	"kratos_frame/internal/server/middleware_http/cors"
	"kratos_frame/internal/server/middleware_http/jwt"
	"kratos_frame/internal/service"
	"net/http"
	"strconv"

	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"

	"github.com/bilibili/kratos/pkg/ecode"

	"kratos_frame/internal/model"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"

	"github.com/bilibili/kratos/pkg/net/http/blademaster/render"
)

var (
	pbService interface{}
	svc       *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine, err error) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err = paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			return
		}
		err = nil
	}
	svc = s
	pbService = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	// 挂载自适应限流中间件到 bm engine，使用默认配置
	limiter := bm.NewRateLimiter(nil)
	e.Use(limiter.Limit())
	// 跨域
	e.Use(cors.Cors())

	api := e.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/start", howToStart)
			user := v1.Group("/user")
			{
				user.POST("/find", jwt.JWT(), FindUser)                                     // 查询所有用户
				user.POST("/info", jwt.JWT(), FindUserInfo)                                 // 个人中心
				user.POST("/signup", Signup)                                                // 注册
				user.POST("/login", Login)                                                  // 登录
				user.POST("/logout", jwt.JWT(), Logout)                                     // 退出登录
				user.POST("/send_sms", SendSms)                                             // 发送短信验证码
				user.POST("/send_email", SendEmail)                                         // 发送邮箱验证码
				user.POST("/forget_password", ForgetPassword)                               // 忘记密码
				user.POST("/update_password", jwt.JWT(), UpdatePassword)                    // 修改密码
				user.POST("/update_fund_password", jwt.JWT(), UpdateFundPassword)           // 修改资金密码
				user.POST("/find_cert_level", jwt.JWT(), FindCertLevel)                     // 查看认证等级
				user.POST("/user_role", jwt.JWT(), InsertRoleUser)                          // 添加用户角色
				user.POST("/invite", jwt.JWT(), FindUserInvite)                             // 查询用户 邀请相关 邀请链接 邀请人数 邀请码
				user.POST("/invite/find_subordinate", jwt.JWT(), FindUserInviteSubordinate) // 查询用户一级下级
				certification := user.Group("/certification")                               // 身份相关
				{
					certification.Use(jwt.JWT())
					certification.POST("/insert", InsertCertification) // 添加身份
					certification.POST("/delete", DeleteCertification) // 删除身份
					certification.POST("/update", UpdateCertification) // 修改身份
					certification.POST("/find", FindCertification)     // 查看身份
				}
			}
			file := v1.Group("/file") // 文件操作
			{
				file.Any("/ueditor", Ueditor)
				file.Use(jwt.JWT())
				file.POST("/upload", FileUpload)
				file.POST("/del", DelFile)
			}
			banner := v1.Group("/banner") // 轮播图
			{
				banner.POST("/find", FindBanner)
				banner.Use(jwt.JWT())
				banner.POST("/insert", InsertBanner)
				banner.POST("/delete", DeleteBanner)
			}
			notice := v1.Group("/notice") //公告
			{
				notice.POST("/find", FindNotice)
				notice.Use(jwt.JWT())
				notice.POST("/insert", InsertNotice)
				notice.POST("/delete", DeleteNotice)
				notice.POST("/update", UpdateNotice)
			}
			version := v1.Group("/version") // 版本更新
			{
				version.POST("/find", FindVersion)
				version.Use(jwt.JWT())
				version.POST("/insert", InsertVersion)
				version.POST("/delete", DeleteVersion)
				version.POST("/update", UpdateVersion)
			}
			protocol := v1.Group("/protocol") // 用户协议
			{
				protocol.POST("/find", FindProtocol)
				protocol.Use(jwt.JWT())
				protocol.POST("/insert", InsertProtocol)
				protocol.POST("/delete", DeleteProtocol)
				protocol.POST("/update", UpdateProtocol)
			}
			work := v1.Group("/work") // 工单
			{
				work.Use(jwt.JWT())
				work.POST("/insert", InsertWork)
				work.POST("/delete", DeleteWork)
				work.POST("/update", UpdateWork)
				work.POST("/find", FindWork)
			}
			menu := v1.Group("/menu") // 菜单管理
			{
				menu.POST("/find", FindMenu)
				menu.Use(jwt.JWT())
				menu.POST("/insert", InsertMenu)
				menu.POST("/delete", DeleteMenu)
				menu.POST("/update", UpdateMenu)
			}
			role := v1.Group("/role") // 角色管理
			{
				role.POST("/find", FindRole)
				role.Use(jwt.JWT())
				role.POST("/insert", InsertRole)
				role.POST("/delete", DeleteRole)
				role.POST("/update", UpdateRole)
				role.POST("/role_menu", InsertRoleMenu) // 角色和菜单绑定
			}
		}
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}

func bindArgs(c *bm.Context, obj interface{}) error {
	if err := c.Bind(&obj); err != nil {
		c.Abort()
		return err
	}
	log.Info(c.Request.URL.Path+"   参数", obj)
	if err := binding.Validator.ValidateStruct(obj); err != nil {
		log.Error("参数验证失败：%s", err.Error())
		c.Abort()
		return err
	}
	return nil
}

// 重写 bm.context.json 使用ttl参数， 用作返回数据总量
func RespJson(c *bm.Context, data interface{}, ttl int, err error) {
	code := http.StatusOK
	c.Error = err
	bcode := ecode.Cause(err)
	// TODO app allow 5xx?
	/*
		if bcode.Code() == -500 {
			code = http.StatusServiceUnavailable
		}
	*/
	header := c.Writer.Header()
	header.Set("kratos-status-code", strconv.FormatInt(int64(bcode.Code()), 10))
	c.Render(code, render.JSON{
		Code:    bcode.Code(),
		Message: bcode.Message(),
		TTL:     ttl,
		Data:    data,
	})
}
