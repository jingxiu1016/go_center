/**
* @file: status.go ==> gateway/handle/app
* @package: app
* @author: jingxiu
* @since: 2022/11/3
* @desc: app.Status()
 */

package app

import "github.com/gin-gonic/gin"

// Status
// @Group[app]
// @Route[status]
// @Method[Post]
// @Middleware[Auth|JWT]
// @Doc[用于更新应用状态]
func (a *App) Status(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
