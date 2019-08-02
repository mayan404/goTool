package mystrconv

import (
	"strconv"
	"strings"
)


/**
字符串转int数组
例如：1-10 -> [1,2,3,4,5,6,7,8,9,10]
 */
func AToIs(str string) []int32 {
	a := strings.Split(str, "-")

	var arr []int32
	if len(a) == 1 {
		i , _ := strconv.Atoi(a[0])
		arr = append(arr,int32(i))

	} else {

		start , _ := strconv.Atoi(a[0])
		end , _ := strconv.Atoi(a[1])
		for i := start; i <=  end ; i++{
			arr = append(arr,int32(i))
		}
	}

	return arr
}