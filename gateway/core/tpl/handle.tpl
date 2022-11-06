/**
* @file: {{.File}}==> {{.Path}}
* @package: router
* @author: jingxiu
* @since: {{.Date}}
* @desc: //Write Doc
 */

package {{.Package}}

import (
    "github.com/gin-gonic/gin"
)

// {{.CurrentHandle}}
// @Group[{{.Package}}]
// @Route[]
// @Method[]
// @Middleware[]
// @Doc[]
func (x *{{.Controller}}) {{.CurrentHandle}}(c *gin.Context) {
	// TODO This is where you should write the logic code
}
