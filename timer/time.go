/*
 * @Date: 2021-06-10 14:41:54
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 14:46:39
 * @FilePath: /yolk/timer/time.go
 */
package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")

	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil
}
