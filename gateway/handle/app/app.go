/**
* @file: app.go ==> gateway/handle/app
* @package: app
* @author: jingxiu
* @since: 2022/11/2
* @desc: 应用创建
 */

package app

import (
	"gateway/services"
	"github.com/gin-gonic/gin"
	"model/dao"
)

type App struct {
	dao *dao.App //数据持久层
}

// Status @Method[Post] @Doc[用于获取实际] @Middleware[Auth|JWT]
func (a *App) Status(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Create @Method[Post] @Doc[用于创建app] @Middleware[Auth|JWT]
func (a *App) Create(c *gin.Context) {
	if err := c.ShouldBind(a.dao); err != nil {
		services.SvcContext.Log.Fatal("应用创建请求数据绑定失败")
		return
	}
}

func (a *App) List(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Get(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *App) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
