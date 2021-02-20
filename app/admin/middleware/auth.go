package middleware

import (
	"time"

	"go-admin/app/admin/middleware/handler"
	jwt "go-admin/pkg/jwtauth"
	"go-admin/tools/config"
)

// AuthInit jwt验证new
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	if config.ApplicationConfig.Mode == "dev" {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if config.JwtConfig.Timeout != 0 {
			timeout = time.Duration(config.JwtConfig.Timeout) * time.Second
		}
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(config.ApplicationConfig.JwtSecret),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler, //身份信息Map
		Authenticator:   handler.Authenticator,   //登陆
		Authorizator:    handler.Authorizator,    //设置上下文件
		Unauthorized:    handler.Unauthorized,    //未经授权 返回的信息
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",  //TokenHead头名称
		TimeFunc:        time.Now,
	})

}
