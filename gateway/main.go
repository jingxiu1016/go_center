/**
* @file: main.go ==> gateway
* @package: gateway
* @author: jingxiu
* @since: 2022/11/2
* @desc: api 网关启动入口
 */

package main

import "gateway/core/jx_api"

func main() {
	// 1. read config from yaml
	//if err := config.ConfigInit("./gateway.yaml"); err != nil {
	//	panic("config file read ersor!")
	//	return
	//}
	// 2. generate route code from handle code comments
	//jx_api.GenRouters()
	jx_api.JingXiu()
	// 3. start some services
	//services.SvcContext = services.NewContext(config.C)
	// 4. start gin router
	//r := router.InitRouter()
	//if err := r.Run(config.C.Gateway.Listen); err != nil {
	//	fmt.Printf("startup service failed, err:%v\n\n", err)
	//}
}
