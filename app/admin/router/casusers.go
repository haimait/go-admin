package router

import (
    "github.com/gin-gonic/gin"

    "go-admin/app/admin/middleware"
    "go-admin/app/admin/models"
    "go-admin/app/admin/service/dto"
    "go-admin/common/actions"
    jwt "go-admin/pkg/jwtauth"
)

func init()  {
	routerCheckRole = append(routerCheckRole, registerCasusersRouter)
}

// 需认证的路由代码
func registerCasusersRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/casusers").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.Casusers{}
        r.GET("", actions.PermissionAction(), actions.IndexAction2(model, new(dto.CasusersSearch), func() interface{} {
            list := make([]models.Casusers, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction2(new(dto.CasusersById), nil))
        r.POST("", actions.CreateAction2(new(dto.CasusersControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction2(new(dto.CasusersControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction2(new(dto.CasusersById)))
    }
}
