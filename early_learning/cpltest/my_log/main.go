package main

import (
	"fmt"
	"log"
	"os"
)

//func main() {
//	log.Println("小魔童打日志 ... ")
//	test := "Hello wrold "
//	// Printf 有格式控制符
//	log.Printf("%s 小魔童打日志 ... \n", test)
//
//	log.Fatalln("小魔童 打日志，触发了 Fatal")
//
//	log.Panicln("小魔童 打日志，触发了 Panic")
//}
//func main() {
//	logger := log.New(os.Stdout, "<XMT>", log.Lshortfile|log.Ldate|log.Ltime)
//	logger.Println("小魔童打印了带有前缀的日志 ... ")
//}

func main() {
	logFile, err := os.OpenFile("./XMT.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("os.OpenFile error :", err)
		return
	}
	// 设置输出位置 ，里面有锁进行控制
	log.SetOutput(logFile)

	// 设置日志属性
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
	// 打印日志
	log.Println("小魔童的 新 日志 ... ")
	// 手动设置前缀
	log.SetPrefix("【重点】")

	log.Println("小魔童的重要日志...")
}
