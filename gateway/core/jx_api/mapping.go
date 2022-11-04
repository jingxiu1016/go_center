/**
* @file: mapping.go ==> core/jx_api
* @package: jx_api
* @author: jingxiu
* @since: 2022/11/2
* @desc: 仅用作接口函数的的映射
 */

package jx_api

//这个设计是为了生成路由接口文件的时候指定接口的指向方法

// APIHandleMapping API方法定向映射
//APIHandleMapping = map[string]func(APIHandler, *gin.Context){
//	"Create": APIHandler.Create,
//	"List":   APIHandler.List,
//	"Info":   APIHandler.Info,
//	"Delete": APIHandler.Info,
//	"Update": APIHandler.Update,
//	"Status": APIHandler.Status,
//}
var (

	// APIMatchMapping API文件扫描匹配字段
	APIMatchMapping = []string{"@Group", "@Route", "@Method", "@Middleware", "@Doc"}
	// APIMiddlewareMapping api路由中间件匹配
	APIMiddlewareMapping = map[string]string{
		"JWT":  "middleware.JWTAuthMiddleware()",
		"Auth": "middleware.UserAuth()",
	}
)
