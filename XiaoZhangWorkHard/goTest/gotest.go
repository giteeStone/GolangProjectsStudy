package split_string

import (
	"strings"
)

// 切割字符串
// example:
// abc,b=>[a c]
func Split(str string, sep string) []string {
	var ret = make([]string, 0, strings.Count(str, sep)+1) //预先分配好内存
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	if str != "" {
		ret = append(ret, str)
	}
	return ret
}
