/**
* @file: {{.filename}}==> {{.filepath}}
* @package: router
* @author: jingxiu
* @since: {{.date}}
* @desc: {{.doc}}
 */

package router

import (
    "gateway/handle/{{.higherDir}}"
    "gateway/middleware"
    "github.com/gin-gonic/gin"
)

func {{.funcName}}(c *gin.Engine){
    instance := {{.higherDir}}.New{{.pak}}Handle()
    group := c.Group("/{{.group}}/")
    { {{range $k,$v := .routers}}
       // {{$v.Doc}}
       group.{{$v.Method}}("{{$v.Route}}",{{$v.Middleware}}, instance.{{$v.Handle}}){{end}}
    }
}
