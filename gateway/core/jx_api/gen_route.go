/**
* @file: gen_route.go ==> core/jx_api
* @package: jx_api
* @author: jingxiu
* @since: 2022/11/2
* @desc: 解析文件，生成
*  这个文件主要负责识别handle文件夹下的所有接口函数，
*  并根据其上特定的注释格式，生成对应的路由文件，
*  生成的路由格式为：r.Post("create",middleware.Auth(),middleware.JWT(), app.Create)
 */

package jx_api

import (
	"bufio"
	"html/template"
	"io"
	"os"
	"strings"
	"time"
)

type GenRoute struct {
	FuncName   string   // 解析匹配到的方法名称
	Group      string   // 解析匹配到的路由组名称
	Route      string   // 解析匹配到的路由名
	Method     string   // 解析匹配到的HTTP方法名
	MiddleWare []string // 解析匹配到的中间件方法
	Doc        string   // 解析文档
}

var handlePath string
var routerPath string
var templatePath string
var infoChan = make(chan []string)
var register = make(map[string][]*GenRoute)

func Gen() {
	workspace, _ := os.Getwd()
	handlePath = workspace + "\\handle"
	routerPath = workspace + "\\router"
	templatePath = workspace + "\\core\\tpl"
	rangeDir(handlePath)
	for key, value := range register {
		writeRouterFile(routerPath, key, value)
	}
}

// 遍历文件夹内容下的内容
func rangeDir(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic("接口目录扫描失败！" + err.Error())
	}
	for _, item := range dir {
		if item.IsDir() {
			rangeDir(handlePath + "\\" + item.Name())
		} else {
			sr := strings.Split(path, "\\")
			register[sr[len(sr)-1]] = append(register[sr[len(sr)-1]], openFile(path+"\\"+item.Name())...)
		}
	}
}

func openFile(file string) []*GenRoute {
	open, err := os.Open(file)
	if err != nil {
		panic(file + "文件打开错误" + err.Error())
	}
	defer open.Close()
	reader := bufio.NewReader(open)
	var results []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(line)
		if strings.Contains(str, "//") && strings.IndexAny(str, "/") == 0 {
			results = append(results, str)
		}
	}
	var routeList = make([]*GenRoute, 0)
	if len(results) > 0 {
		routeList = append(routeList, matchKeywords(results))
	}
	return routeList
}

// 匹配注释中的关键词
func matchKeywords(info []string) *GenRoute {
	first := trimPrefix(info[0])
	temp := &GenRoute{
		FuncName: first,
	}
	for _, item := range APIMatchMapping {
		for _, fo := range info[1:] {
			if strings.Contains(fo, item) {
				switch item {
				case "@Group":
					group := trimPrefix(fo)
					left, right := indexBrackets(group)
					temp.Group = group[left+1 : right]
				case "@Route":
					route := trimPrefix(fo)
					left, right := indexBrackets(route)
					temp.Route = route[left+1 : right]
				case "@Method":
					method := trimPrefix(fo)
					left, right := indexBrackets(method)
					temp.Method = method[left+1 : right]
				case "@Middleware":
					mw := trimPrefix(fo)
					left, right := indexBrackets(mw)
					temp.MiddleWare = strings.Split(mw[left+1:right], "|")
				case "@Doc":
					doc := trimPrefix(fo)
					left, right := indexBrackets(doc)
					temp.Doc = doc[left+1 : right]
				}
			}
		}
	}
	return temp
}

// 写入路由文件
func writeRouterFile(path, key string, value []*GenRoute) {
	filename := key + "_router.gen.go"
	funcName := firstUpper(key) + "Router"
	//写入文件时，使用带缓存的 *Writer
	data := map[string]interface{}{
		"filename":   filename,
		"filepath":   path,
		"date":       time.Now().Format("2016/01/01"),
		"doc":        key + " 路由",
		"funcName":   funcName,
		"higherDir":  key,
		"pak":        firstUpper(key),
		"group":      value[0].Group,
		"httpMethod": strings.ToUpper(value[0].Method),
		"route":      value[0].Route,
		"middleware": transitMiddle(value[0].MiddleWare),
		"handle":     value[0].FuncName,
	}
	tmp := template.Must(template.ParseFiles(templatePath + "\\route.tpl"))
	create, err := os.OpenFile(path+"\\"+filename, os.O_CREATE, 0666)
	if err != nil {
		panic(filename + "文件创建失败")
	}
	defer create.Close()
	err = tmp.Execute(create, data)
	if err != nil {
		panic(filename + " 模板文件生成失败：" + err.Error())
	}
}

func trimPrefix(s string) string {
	s = strings.TrimPrefix(s, "//")
	s = strings.TrimSpace(s)
	return s
}
func indexBrackets(s string) (int, int) {
	return strings.Index(s, "["), strings.Index(s, "]")
}

func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func transitMiddle(mi []string) string {
	str := ""
	for _, item := range mi {
		str += APIMiddlewareMapping[item] + ", "
	}
	return str[:len(str)-2]
}
