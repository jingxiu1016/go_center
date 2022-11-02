/**
* @file: gen.go ==> model/cmd
* @package: model
* @author: jingxiu
* @since: 2022/11/1
* @desc: 使用gorm-gen快速生成持久化数据层与逻辑层通信层
 */

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./query",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./dao",
	})
	g.UseDB(ConnectDB())
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

func ConnectDB() (conn *gorm.DB) {
	conn, err := gorm.Open(mysql.Open("center:center@tcp(43.138.235.141:3306)/center?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
