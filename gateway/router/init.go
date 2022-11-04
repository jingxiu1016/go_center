/**
* @file: init.go ==> gateway/router
* @package: router
* @author: jingxiu
* @since: 2022/11/2
* @desc: //TODO
 */

package router

import (
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter init router
func InitRouter() *gin.Engine {
	r := initGin()
	//r.Use(middleware.ParseIP())
	loadRoute(r)
	return r
}

// init Gin
func initGin() *gin.Engine {
	//设置gin模式
	//gin.SetMode(global.VP.GetString("RunMode"))
	engine := gin.New()
	// 使用日志打印
	engine.Use(gin.Logger())
	//定义404中间件
	engine.NoRoute(middleware.NoRoute())
	return engine
}

// 加载路由
func loadRoute(r *gin.Engine) {
	//UserRoute(r)
	//AppRouter(r)
	//for _, handler := range regList {
	//	switch handler.method {
	//	case http.MethodGet:
	//		r.GET(handler.path, handler.handles...)
	//	case http.MethodPost:
	//		r.POST(handler.path, handler.handles...)
	//	case http.MethodPut:
	//		r.PUT(handler.path, handler.handles...)
	//	case http.MethodPatch:
	//		r.PATCH(handler.path, handler.handles...)
	//	case http.MethodDelete:
	//		r.DELETE(handler.path, handler.handles...)
	//	}
	//}
}
