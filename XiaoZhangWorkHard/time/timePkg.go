// time包
// time.Time类型表示时间。time.Now()函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息
// 一个很重要的时间：2006-01-02 15：04：05.000
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳。
	fmt.Println(now.Unix()) //Unix（）将时间戳转化为时间格式
	fmt.Println(now.UnixNano())

	ret := time.Unix(1649056537, 0) //将时间戳转化为时间
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	//时间间隔
	//time包中定义的时间间隔常量如下：
	// const{
	// 	Nanosecond Duration=1
	// 	Microsecond =1000*Nanosecond
	// 	Millisecond =1000*Microsecond
	// 	Second      =1000*Millisecond
	// 	Minute      =60*Second
	// 	Hour        =60*Minute
	// }
	//time包中还有Add、Sub来进行时间运算
	//time包中还有Equal来判断两个时间是否相等
	//time包中还有Before和After来判断谁先谁后，func (t Time) Before(u Time) bool

	fmt.Println(time.Second)
	fmt.Println(now.Add(24 * time.Hour)) //加一天

	//定时器
	//timer := time.Tick(time.Second) //一秒钟给一个时间值
	// for t := range timer {
	// 	fmt.Println(t) //1秒钟执行一次
	// }

	//时间格式化
	fmt.Println(now.Format("2006-01-02"))              //将现在的时间转化为xxxx-xx-xx类型输出
	fmt.Println(now.Format("2006/01/02 15:04:05"))     //将现在的时间转化为xxxx/xx/xx xx:xx:xx类型输出
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))  //以12进制 PM类型输出
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) //打印毫秒

	//按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-08-03") //2000
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}
