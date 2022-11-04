/**
* @file: app.go ==> gateway/handle/app
* @package: app
* @author: jingxiu
* @since: 2022/11/2
* @desc: 应用创建
 */

package app

import (
	"model/dao"
)

type App struct {
	dao *dao.App //数据持久层
}

func NewAppHandle() *App {
	return &App{}
}
