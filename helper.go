/*
 * @Date: 2021-03-12 17:37:03
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-26 21:18:37
 * @FilePath: /yolk/helper.go
 */
package yolk

import "strconv"

func ToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case uint:
		return strconv.Itoa(int(i.(uint)))
	}
	return ""
}
