/**
* @file: main.go ==> gateway
* @package: gateway
* @author: jingxiu
* @since: 2022/11/2
* @desc: api 网关启动入口
 */

package main

import (
	"fmt"
	"gateway/config"
	"gateway/router"
	"gateway/services"
)

func main() {
	// 1. read config from yaml
	if err := config.ConfigInit("etc/gateway.yaml"); err != nil {
		panic("config file read ersor!")
		return
	}
	services.SvcContext = services.NewContext(config.C)
	r := router.InitRouter()
	if err := r.Run(config.C.Gateway.Listen); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
