package golang_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 获取当前的时间区设置

	// 当前时间的状态
	dateTime := time.Now()
	fmt.Println(dateTime)

	now := time.Now()
	local1, err1 := time.LoadLocation("") //等同于"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 := time.LoadLocation("Local") //服务器设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 := time.LoadLocation("America/Los_Angeles")
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))

	// 设置时区
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	fmt.Println("SH : ", time.Now().In(cstZone).Format("2006-01-02 15:04:05"))

	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //  上海
	fmt.Println(time.Now().In(cstSh).Format("2006-01-02 15:04:05"))

	// 当前的日期
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// 日期的格式化
	formate := "2006-01-02 15:04:05 Mon"
	local1, err1 = time.LoadLocation("UTC") //输入参数"UTC"，等同于""
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 = time.LoadLocation("Local")
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 = time.LoadLocation("America/Los_Angeles")
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(now.In(local1).Format(formate))
	fmt.Println(now.In(local2).Format(formate))
	fmt.Println(now.In(local3).Format(formate))

	// 日期初始化
	local, _ := time.LoadLocation("America/Los_Angeles")
	timeFormat := "2006-01-02 15:04:05"
	//func Unix(sec int64, nsec int64) Time {
	time1 := time.Unix(1480390585, 0)                                          //通过unix标准时间的秒，纳秒设置时间
	time2, _ := time.ParseInLocation(timeFormat, "2016-11-28 19:36:25", local) //洛杉矶时间
	fmt.Println(time1.In(local).Format(timeFormat))
	fmt.Println(time2.In(local).Format(timeFormat))
	chinaLocal, _ := time.LoadLocation("Local") //运行时，该服务器必须设置为中国时区，否则最好是采用"Asia/Chongqing"之类具体的参数。
	fmt.Println(time2.In(chinaLocal).Format(timeFormat))

	// 当前的时间戳（秒、毫秒）
	timeUnix := time.Now().Unix()         //单位秒
	timeUnixNano := time.Now().UnixNano() //单位纳秒

	fmt.Println(timeUnix)
	fmt.Println(timeUnixNano)

	// 日期的运算

	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1)

	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)

	// printSplit(50)

	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)

	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	// printSplit(50)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)

	// 获取今天0点0时0分的时间戳
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	fmt.Println(startTime)
	fmt.Println(startTime.Format("2006/01/02 15:04:05"))

	// 02: 获取今天23:59:59秒的时间戳
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	fmt.Println(endTime)
	fmt.Println(endTime.Format("2006/01/02 15:04:05"))

	// 08: 判断一个时间是否在一个时间之后
	stringTime, _ := time.Parse("2006-01-02 15:04:05", "2019-12-12 12:00:00")
	beforeOrAfter := stringTime.After(time.Now())

	if true == beforeOrAfter {
		fmt.Println("2019-12-12 12:00:00在当前时间之后!")
	} else {
		fmt.Println("2019-12-12 12:00:00在当前时间之前!")
	}

	// 09: 判断一个时间相比另外一个时间过去了多久

	startTime = time.Now()
	time.Sleep(time.Second * 5)

	fmt.Println("离现在过去了：", time.Since(startTime))

	// 日期时间戳间的转换

	// ref doc
	// https: //segmentfault.com/a/1190000019694913?utm_source=sf-similar-article
}

func TestStringToStamp(t *testing.T) {
	// tmpH1Time := fmt.Sprintf("%s/%s", "11/30", "2021")

	// fmt.Println(tmpH1Time)
	// tm2, _ := time.Parse("01/02/2006", tmpH1Time)
	// fmt.Println(int(tm2.Unix()))

	// sTime := fmt.Sprintf("%s-0%s-01 00:00:00", "2021", "2")

	// tm2, _ := time.Parse("2006-01-02 15:04:05", sTime)
	// fmt.Println(int(tm2.Unix() - 8*60*60 + 60*60*24*30))

	// year, month, _ := time.Now().Date()
	// curMonth := fmt.Sprintf("%d%d", year, int(month))

	// fmt.Println(year)
	// fmt.Println(int(month))

	// eTime := fmt.Sprintf("%d-12-31 23:59:59", 2021)
	// tm2, _ := time.Parse("2006-01-02 15:04:05", eTime)
	// endTime := int(tm2.Unix() - 8*60*60)
	// fmt.Println(int(endTime))

	// 日期格式
	// 时间戳

	// local, _ := time.LoadLocation("Asia/Shanghai")

	// timeFormat := "2006-01-02 15:04:05"
	// time2, _ := time.ParseInLocation(timeFormat, "2021-05-22 00:00:00", local) //洛杉矶时间

	// fmt.Println(time2.Unix())
	now := time.Now()
	// m, _ := time.ParseDuration("-1m")
	// m1 := now.Add(m)
	// fmt.Println(m1)

	m, _ := time.ParseDuration("24h")
	result := now.Add(m)
	fmt.Println(result)
	fmt.Println(result.Format("2006/01/02 15:04:05"))

}
