/**
* @file: list.go ==> gateway/handle/app
* @package: app
* @author: jingxiu
* @since: 2022/11/3
* @desc: app.List()
 */

package app

import "github.com/gin-gonic/gin"

// List
// @Group[app]
// @Route[list]
// @Method[get]
// @Middleware[Auth|JWT]
// @Doc[获取应用列表]
func (a *App) List(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
