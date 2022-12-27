package main

import (
	"fmt"
	"time"
)

func timeShow() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//time.Unix()
	ret := time.Unix(1658671509, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	//时间间隔
	fmt.Println(time.Second)
	fmt.Println(now.Add(24 * time.Hour))

	//时间的格式化
	//12 hours
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	//24 hours
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 03:04:05.000"))
	//定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }
	nextYear, err := time.Parse("2006-01-02 03:04:05", "2022-07-24 00:00:00")
	if err != nil {
		fmt.Printf("Parse time failed, err:%v\n", err)
		return
	}
	nextYear = nextYear.UTC()
	d := nextYear.Sub(now)
	fmt.Println(d)

}

func main() {
	timeShow()
}
