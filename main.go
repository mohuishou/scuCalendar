package main

import (
	"fmt"
	"os"

	"log"

	"github.com/mohuishou/scujwc-go"
)

func main() {
	fmt.Println("******************************************************")
	fmt.Println("欢迎使用四川大学课程表转日历文件工具")
	fmt.Println("纯本地应用，不会上传任何信息到私人服务器")
	fmt.Println("代码开源,欢迎Star：https://github.com/mohuishou/scuCalendar.git")
	fmt.Println("微信搜索订阅号：四川大学飞扬俱乐部，报修/绩点/IT资讯等你来查")
	fmt.Println("微信搜索服务号：scuplus，更多小工具不定期推出")
	fmt.Println("Power By Mohuishou of FYSCU")
	fmt.Println("******************************************************")
	fmt.Println("")

	// 登录
	var j scujwc.Jwc
	var (
		uid      int
		password string
	)

	var ical scujwc.Ical
	for {
		fmt.Println("请输入您的学号：")
		fmt.Scan(&uid)
		fmt.Println("请输入您的密码：")
		fmt.Scan(&password)
		log.Println("教务处登录中...")

		err := j.Init(uid, password)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("教务处登录成功,正在获取课程数据...")
		ical, err = j.Calendar()
		if err != nil {
			log.Println(err)
			continue
		}
		break
	}
	log.Println("课程数据获取成功，正在生成文件...")
	f, err := os.Create("ical.ics")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(ical.Bytes())
	log.Println("文件生成完成，文件名：ical.ics")

	fmt.Println()
	fmt.Println("**********************输入任意字符回车退出*********************")
	fmt.Println("微信搜索订阅号：四川大学飞扬俱乐部，报修/绩点/IT资讯等你来查")
	fmt.Println("微信搜索服务号：scuplus，更多小工具不定期推出")
	fmt.Println("**********************输入任意字符回车退出*********************")
	fmt.Scan(&password)
}
