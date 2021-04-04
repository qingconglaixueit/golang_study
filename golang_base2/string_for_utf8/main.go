package main

import "fmt"

//string 本质上是一个byte数组
func eng() {
	str := "hellowolrd"
	fmt.Println("len == ",len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ",str[i])
	}
	fmt.Println()
}

func ch(){
	str := "你好world"
	fmt.Println("len == ",len(str))
	for index,v := range str{
		fmt.Printf("index = %d,v = %x\n",index,v)
	}

}

//遍历中文 和英文字符串 ,在go里面中文默认是utf-8的格式，默认在三个字节
//len ==  10
//68 65 6c 6c 6f 77 6f 6c 72 64
//==================
//len ==  11
//index = 0,v = 4f60
//index = 3,v = 597d
//index = 6,v = 77
//index = 7,v = 6f
//index = 8,v = 72
//index = 9,v = 6c
//index = 10,v = 64
func main() {

	eng()

	fmt.Println("==================")

	ch()

}
