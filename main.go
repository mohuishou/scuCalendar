package main

import (
	"fmt"
	"os"

	"log"

	"github.com/mohuishou/scujwc-go"
)

func main() {
	fmt.Println("******************************************************\n ")
	fmt.Println("欢迎使用四川大学课程表转日历文件工具 \n ")
	fmt.Println("纯本地应用，不会上传任何信息到私人服务器 \n ")
	fmt.Println("代码开源,欢迎Star：https://github.com/mohuishou/scuCalendar.git \n ")
	fmt.Println("微信搜索订阅号：四川大学飞扬俱乐部，报修/绩点/IT资讯等你来查 \n ")
	fmt.Println("微信搜索服务号：scuplus，更多小工具不定期推出 \n ")
	fmt.Println("Power By Mohuishou of FYSCU \n ")
	fmt.Println("******************************************************")
	fmt.Println("")

	// 登录
	var (
		uid      int
		campus   int
		password string
	)

	var ical scujwc.Ical
	for {
		fmt.Println("请输入您的学号：")
		fmt.Scan(&uid)
		fmt.Println("请输入您的密码：")
		fmt.Scan(&password)
		fmt.Println("请输入您的校区（江安校区请输入1，望江/华西校区:输入2）：")
		fmt.Scan(&campus)

		if campus < 1 || campus > 2 {
			log.Println("请输入正确的校区代码：1 or 2")
			continue
		}

		log.Println("教务处登录中...")

		j, err := scujwc.NewJwc(uid, password)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("教务处登录成功,正在获取课程数据...")
		ical, err = j.Calendar(campus)
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
	fmt.Println("**********************输入任意字符回车退出********************* \n ")
	fmt.Println("微信搜索订阅号：四川大学飞扬俱乐部，报修/绩点/IT资讯等你来查 \n ")
	fmt.Println("微信搜索服务号：scuplus，更多小工具不定期推出 \n ")
	fmt.Println("**********************输入任意字符回车退出*********************")
	fmt.Scan(&password)
}
