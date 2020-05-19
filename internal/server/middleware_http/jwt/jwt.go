package jwt

import (
	"context"
	err_code "kratos_frame/internal/pkg/ecode"
	pkg_jwt "kratos_frame/internal/pkg/jwt"
	"strconv"
	"time"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/container/pool"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/dgrijalva/jwt-go"
	"github.com/siddontang/go/log"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func JWT() bm.HandlerFunc {
	return func(c *bm.Context) {

		// 需要加 判断操作人是不是本人的操作
		b := true
		// 先判断token是否是对的， 是否过期
		Authorization := c.Request.Header.Get("Authorization")
		if Authorization == "" {
			b = false
		} else {
			claims, err := pkg_jwt.ParseToken(Authorization)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					b = false
				default:
					b = false
				}
			}
			if b && claims != nil {
				// 再判断此用户的token 是否是最新的（每次登陆都会更新token）
				r := redis.NewRedis(&redis.Config{
					Config: &pool.Config{
						Active:      10,
						Idle:        10,
						IdleTimeout: xtime.Duration(10 * time.Second),
					},
					Name:         "kratos_frame",
					Proto:        "tcp",
					Addr:         "127.0.0.1:6379",
					DialTimeout:  xtime.Duration(5 * time.Second),
					ReadTimeout:  xtime.Duration(5 * time.Second),
					WriteTimeout: xtime.Duration(5 * time.Second),
				}, redis.DialPassword("asia@123!"))
				log.Info(strconv.Itoa(int(claims.Id)) + "_token")
				token, err := redis.String(r.Do(context.Background(), "GET", strconv.Itoa(int(claims.Id))+"_token"))
				if err != nil && err.Error() != "redigo: nil returned" {
					b = false
				}
				if token != "" && Authorization != token {
					b = false
				}
				defer r.Close()
			}
		}

		if !b {
			c.JSON(nil, err_code.VerifyTokenErr)
			c.Abort()
			return
		}
		c.Next()
	}
}
