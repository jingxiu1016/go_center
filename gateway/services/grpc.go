/**
* @file: grpc.go ==> gateway/services
* @package: services
* @author: jingxiu
* @since: 2022/11/2
* @desc: grpc 链接器
 */

package services

import "gateway/config"

// GrpcContext 填入gRpc客户端注册实例
type GrpcContext struct {
}

// GrpcInit 初始化gRpc链接
func GrpcInit(c *config.Config) *GrpcContext {
	var g GrpcContext

	return &g
}
