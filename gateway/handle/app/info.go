/**
* @file: info.go ==> gateway/handle/app
* @package: app
* @author: jingxiu
* @since: 2022/11/3
* @desc: app.Info()
 */

package app

import "github.com/gin-gonic/gin"

// Info
// @Group[app]
// @Route[info]
// @Method[get]
// @Middleware[Auth|JWT]
// @Doc[获取单个应用的详情]
func (a *App) Info(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
