package permission

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
	"kratos_frame/internal/service"
	"strconv"

	"github.com/bilibili/kratos/pkg/log"

	pkg_jwt "kratos_frame/internal/pkg/jwt"

	"github.com/dgrijalva/jwt-go"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func CasbinMiddleware(svc *service.Service) bm.HandlerFunc {
	return func(c *bm.Context) {
		Authorization := c.Request.Header.Get("Authorization")
		t, _ := jwt.Parse(Authorization, func(*jwt.Token) (interface{}, error) {
			return pkg_jwt.JwtSecret, nil
		})
		log.Info("", pkg_jwt.GetIdFromClaims("id", t.Claims), c.Request.URL.Path, c.Request.Method)

		// 1. 获取用户角色
		id, _ := strconv.Atoi(pkg_jwt.GetIdFromClaims("id", t.Claims))
		roles, err := svc.FindRoleByUser(c.Context, model.FindRoleByUserParams{UserId: uint(id)})
		if err.Code() != 200 {
			c.JSON(nil, err)
			c.Abort()
			return
		}
		var b bool
		// 2. 根据请求路径获取(菜单)mid
		menus, err := svc.FindMenuByPath(c.Context, model.FindMenuByPathParams{Path: c.Request.URL.Path})
		if err.Code() != 200 {
			c.JSON(nil, err)
			c.Abort()
			return
		}
		if len(menus) == 0 {
			c.JSON(nil, err_code.VerifyRoleErr)
			c.Abort()
			return
		}
		// 3. 对比判断是否有权限
		for _, v := range roles {
			e, err := svc.EnforcePolicy(model.EnforcePolicyParams{
				Name:   strconv.Itoa(int(v.RoleId)),
				Path:   menus[0].Mid,
				Method: c.Request.Method,
			})
			if err.Code() != 200 {
				c.JSON(nil, err)
				c.Abort()
				return
			}
			if e {
				b = true
				break
			}
		}
		if !b {
			c.JSON(nil, err_code.VerifyRoleErr)
			c.Abort()
			return
		}
		c.Next()
	}
}
