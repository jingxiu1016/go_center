/**
* @file: generic.go ==> common/utils
* @package: utils
* @author: jingxiu
* @since: 2022/11/5
* @desc: 一些泛型使用
 */

package utils

// Reverse 反转切片
func Reverse[T string | int | int32 | int64](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
