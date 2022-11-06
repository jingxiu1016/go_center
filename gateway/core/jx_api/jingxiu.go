/**
* @file: jingxiu.go ==> gateway/core/jx_api
* @package: jx_api
* @author: jingxiu
* @since: 2022/11/5
* @desc: 命令行工具启动入口
 */

package jx_api

import (
	"common/utils"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"os"
)

// JingXiu 主命令
func JingXiu() {
	app := &cli.App{
		Name:  "jingxiu",
		Usage: "你好，这是一个代码快速生成命令",
		Commands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "创建应用控制器",
				Action:  createController,
			},
			{
				Name:    "route",
				Aliases: []string{"r"},
				Usage:   "生成路由组",
				Action:  generateRouters,
			},
			{
				Name:    "model",
				Aliases: []string{"m"},
				Usage:   "在当前目录下，从配置的链接数据库中生成dao层",
				Action:  generateDatabase,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "src",
						Aliases:  []string{"s"},
						Usage:    "指定要读取的配置文件",
						Required: true,
					},
				},
			},
		},
		Flags: []cli.Flag{},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal("命令未启动：" + err.Error())
	}
}

// createController 定义生成应用控制器命令：
// jingxiu create {$app} {$APIHandler}
// $app 应用名称 $APIHandler 实现的接口
func createController(c *cli.Context) error {
	args := c.Args()
	var argsSlice []string
	if _, ok := APIHandleMapping[args.First()]; ok {
		argsSlice = utils.Reverse(args.Slice())
		fmt.Println("温馨提示：其实我并不建议你先写要实现的接口，你应该先将控制器明确...")
	} else {
		argsSlice = args.Slice()
	}
	GenHandles(argsSlice)
	return nil
}

// generateRouters 生成路由
// jingxiu route [append user]
func generateRouters(c *cli.Context) error {
	args := c.Args()
	workspace, _ := os.Getwd()
	handlePath = workspace + "\\handle"
	routerPath = workspace + "\\router"
	templatePath = workspace + "\\core\\tpl"
	if args.Len() <= 0 {
		rangeDir(handlePath)
		for key, value := range register {
			writeRouterFile(routerPath, key, value)
		}
	} else {
		if args.Slice()[0] == "append" {
			for _, item := range args.Slice()[1:] {
				rangeDir(handlePath + "\\" + item)
				for key, value := range register {
					writeRouterFile(routerPath, key, value)
				}
			}
		} else {
			fmt.Println("如果生成路由文件需要指定某个文件的时候，那么应该使用 append 命令")
		}
	}
	return nil
}

// generateDatabase 生成数据库模型，会在当前目录下
// Jingxiu model -src gateway/gateway.yaml
// "
func generateDatabase(c *cli.Context) error {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./query",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./model",
	})
	path := c.String("src")
	conn, err := ConnectDB(path)
	if err != nil {
		fmt.Printf("%#v", err.Error())
		return err
	}
	g.UseDB(conn)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
	return nil
}

func ConnectDB(path string) (conn *gorm.DB, err error) {
	type DB struct {
		Type   string `yaml:"Type"`   // 链接类型
		Source string `yaml:"Source"` // 链接dns地址
	}
	config := &struct {
		DB DB `yaml:"DB"`
	}{}
	if f, err := os.Open(path); err != nil {
		return nil, errors.New("配置文件读取失败：" + err.Error())
	} else {
		err := yaml.NewDecoder(f).Decode(config)
		if err != nil {
			return nil, errors.New("配置文件读取失败：" + err.Error())
		}
	}
	conn, err = gorm.Open(mysql.Open(config.DB.Source), &gorm.Config{})
	if err != nil {
		//panic(fmt.Errorf("cannot establish db connection: %w", err))
		return nil, errors.New("数据库链接失败：" + err.Error())
	}
	return conn, nil
}
