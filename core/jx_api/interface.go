/**
* @file: interface.go ==> core/jx_api
* @package: jx_api
* @author: jingxiu
* @since: 2022/11/2
* @desc: 网关层一些公用的抽象接口
 */

package jx_api

type APIHandler interface {
	Create()
	List()
	Info()
	Delete()
	Update()
	Status()
}
