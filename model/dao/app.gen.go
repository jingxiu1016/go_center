// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"crypto/md5"
	"fmt"
	"time"
)

const TableNameApp = "app"

// App mapped from table <app>
type App struct {
	ID          int32     `gorm:"column:id;primaryKey" json:"id"`                                             // 主键
	Name        string    `gorm:"column:name;not null" json:"name"`                                           // 应用名称
	Describe    string    `gorm:"column:describe;not null" json:"describe"`                                   // 描述
	Domain      int32     `gorm:"column:domain;not null" json:"domain"`                                       // 领域范畴
	Key         string    `gorm:"column:key;not null" json:"key"`                                             // 应用key
	Secret      string    `gorm:"column:secret;not null" json:"secret"`                                       // 应用密钥
	Status      int32     `gorm:"column:status;not null" json:"status"`                                       // 应用状态： 0 未激活; 1 未使用；2 使用中；3 已销毁；4 阻断中；5 限制中
	CreatedTime time.Time `gorm:"column:created_time;not null;default:CURRENT_TIMESTAMP" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"column:updated_time;not null;default:CURRENT_TIMESTAMP" json:"updated_time"` // 更新时间
}

// TableName App's table name
func (*App) TableName() string {
	return TableNameApp
}

func (a *App) Encrypt(salt string) {
	if a == nil || a.Key == "" || salt == "" {
		panic("empty app")
	}
	str := fmt.Sprintf("%v%v", a.Key, salt)
	a.Secret = fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func (a *App) Verify(salt string) bool {
	if a == nil || a.Key == "" || salt == "" {
		panic("empty app")
	}
	str := fmt.Sprintf("%v%v", a.Key, salt)
	if a.Secret == fmt.Sprintf("%x", md5.Sum([]byte(str))) {
		return true
	}
	return false
}

// 给字符串生成md5
// @params str 需要加密的字符串
// @params salt interface{} 加密的盐
// @return str 返回md5码
//func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
//	if l := len(salt); l > 0 {
//		slice := make([]string, l+1)
//		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
//	}
//	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
//}
