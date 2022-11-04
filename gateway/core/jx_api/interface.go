/**
* @file: interface.go ==> core/jx_api
* @package: jx_api
* @author: jingxiu
* @since: 2022/11/2
* @desc: 网关层一些公用的抽象接口
 */

package jx_api

import "github.com/gin-gonic/gin"

type APIHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Info(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Status(c *gin.Context)
}
