// Code generated by hertz generator.

package main

import (
	handler "github.com/Tracy-coder/e-mall/biz/handler"
	"github.com/Tracy-coder/e-mall/biz/handler/middleware"
	"github.com/Tracy-coder/e-mall/configs"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.POST("/api/v1/user/login", middleware.GetJWTMiddleware(configs.Data(), data.Default()).LoginHandler)
	// your code ...
}
