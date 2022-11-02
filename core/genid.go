/**
* @file: genid.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/2
* @desc: 生成唯一 id
 */

package core

import (
	"common/utils"
	uuid "github.com/satori/go.uuid"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"
)

func GetUUID() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func GetIntNumber(l int) int64 {
	if l > 20 {
		panic("GetIntNumber err")
	}
	num, _ := strconv.ParseInt(GetNumber(l), 10, 64)
	return num
}

func GetNumber(l int) string {
	str := "123456789"
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	return GetOrderId(string(str[r.Intn(len(str))]), l)
}

func GetOrderId(head string, l int) string {
	str := "0123456789"
	result := []byte(head)
	var temp = GetUUID()
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l-len(head); i++ {
		t := temp[r.Intn(len(temp)-1)] % 10
		result = append(result, str[t])
	}
	return string(result)
}

func GetDateId(l int) string {
	return GetOrderId(utils.ToyyyyMMddHHmmss(), l)
}

func GetDateIdInt(l int) int64 {
	str := GetDateId(l)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
