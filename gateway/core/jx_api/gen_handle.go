/**
* @file: gen_handle.go ==> gateway/core/jx_api
* @package: jx_api
* @author: jingxiu
* @since: 2022/11/4
* @desc: //TODO
 */

package jx_api

import (
	"fmt"
	"html/template"
	"os"
	"os/user"
	"strings"
	"time"
)

type GenController struct {
	File          string   // 要生成的文件名
	Path          string   // 要生成的文件路径
	User          string   // 当前的用户昵称
	Date          string   // 当前的日期
	Package       string   // 要定义的包
	Controller    string   // 要生成的控制器
	Interface     string   // 要生成实现的接口
	Handle        []string // 要生成的所有方法名称
	CurrentHandle string   // 当前要生成的方法
}

func GenHandles(s []string) {
	fmt.Println(s)
	workspace, err := os.Getwd()
	if err != nil {
		panic("路径获取错误")
	}
	templatePath = workspace + "\\core\\tpl"
	u, _ := user.Current()
	gen := &GenController{
		File:       s[0] + ".go",
		Path:       workspace + "\\handle\\" + s[0],
		User:       u.Name,
		Date:       time.Now().Format("01/02/2006"),
		Package:    s[0],
		Controller: firstUpper(s[0]),
		Interface:  s[1],
		Handle:     APIHandleMapping[s[1]],
	}
	// 根据模板生成文件，先根据控制器生成目录
	err = os.Mkdir(gen.Path, os.ModePerm)
	if err != nil {
		panic(gen.Path + "文件夹生成错误: " + err.Error())
	}
	//	在控制器目录下生成控制器文件
	ctl, err := os.Create(gen.Path + "\\" + gen.File)
	if err != nil {
		panic(gen.Path + gen.File + "文件夹生成错误: " + err.Error())
	}
	defer ctl.Close()
	ctltmp := template.Must(template.ParseFiles(templatePath + "\\controller.tpl"))
	if err = ctltmp.Execute(ctl, gen); err != nil {
		panic(ctl.Name() + " 模板文件生成失败：" + err.Error())
	}
	handletmp := template.Must(template.ParseFiles(templatePath + "\\handle.tpl"))
	for _, item := range gen.Handle {
		filename := strings.ToLower(item) + ".go"
		currentHandle(filename, gen, item, handletmp)
	}
}

func currentHandle(filename string, gen *GenController, item string, tmp *template.Template) {
	file, err := os.Create(gen.Path + "\\" + filename)
	if err != nil {
		panic(gen.Path + filename + "文件生成错误: " + err.Error())
	}
	defer file.Close()
	gen.CurrentHandle = item
	if err = tmp.Execute(file, gen); err != nil {
		panic(file.Name() + "模板文件生成失败: " + err.Error())
	}
}
