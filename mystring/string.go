package mystring

import (
	"strconv"
	"strings"
)

/**
字符串补齐方法
input 		待补齐的字符串
pad_length 	补齐后的长度
pad_string 	补齐的内容
例如：StrPad("1",3,"0")  -> 001
 */
func StrPad(input string,pad_length int,pad_string string) string  {

	length := pad_length - len(input)
	if length <= 0 {
		return input
	}

	var str = ""
	for i := 0; i < length; i++ {
		str += pad_string
	}

	return str + input
}


