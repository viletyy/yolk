/*
 * @Date: 2021-06-10 14:24:22
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 14:24:36
 * @FilePath: /yolk/random/random.go
 */
package random

import "crypto/rand"

func RandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
