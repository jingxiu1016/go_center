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
	ctx *gin.Context
}

func NewAppHandle(c *gin.Context) *App {
	return &App{
		ctx: c,
	}
}

// Status
// @Group[app]
// @Route[status]
// @Method[Post]
// @Middleware[Auth|JWT]
// @Doc[用于更新应用状态]
func (a *App) Status() {
	//TODO implement me
	panic("implement me")
}

// Create
// @Group[app]
// @Route[create]
// @Method[Post]
// @Middleware[Auth|JWT]
// @Doc[用于更新应用状态]
func (a *App) Create() {
	if err := a.ctx.ShouldBind(a.dao); err != nil {
		services.SvcContext.Log.Fatal("应用创建请求数据绑定失败")
		return
	}
}

// List
// @Group[app]
// @Route[list]
// @Method[get]
// @Middleware[Auth|JWT]
// @Doc[获取应用列表]
func (a *App) List() {
	//TODO implement me
	panic("implement me")
}

// Info
// @Group[app]
// @Route[info]
// @Method[get]
// @Middleware[Auth|JWT]
// @Doc[获取应用列表的详情]
func (a *App) Info() {
	//TODO implement me
	panic("implement me")
}

// Delete
// @Group[app]
// @Route[delete]
// @Method[delete]
// @Middleware[Auth|JWT]
// @Doc[删除应用]
func (a *App) Delete() {
	//TODO implement me
	panic("implement me")
}

// Update
// @Group[app]
// @Route[update]
// @Method[post]
// @Middleware[Auth|JWT]
// @Doc[更新应用详情]
func (a *App) Update() {
	//TODO implement me
	panic("implement me")
}
