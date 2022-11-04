/**
* @file: create.go ==> gateway/handle/app
* @package: app
* @author: jingxiu
* @since: 2022/11/3
* @desc: app.Create()
 */

package app

import (
	"gateway/services"
	"github.com/gin-gonic/gin"
)

// Create
// @Group[app]
// @Route[create]
// @Method[Post]
// @Middleware[Auth|JWT]
// @Doc[用于创建应用]
func (a *App) Create(c *gin.Context) {
	if err := c.ShouldBind(a.dao); err != nil {
		services.SvcContext.Log.Fatal("应用创建请求数据绑定失败")
		return
	}
}
