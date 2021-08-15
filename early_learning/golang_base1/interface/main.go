package main

import "fmt"

//接口
type phone interface {
	call()
	show()
}

type xiaomi struct {
	name string
	ads  string
}

type huawei struct {
	name string
	ads  string
}

//接口实现
func (x xiaomi) call() {
	fmt.Println("phoneName :", x.name)
}

func (x xiaomi) show() {
	fmt.Println("advertisement :", x.ads)
}

func (h huawei) call() {
	fmt.Println("phoneName :", h.name)
}

func (h huawei) show() {
	fmt.Println("advertisement :", h.ads)
}

func main() {
	x := xiaomi{"mi note2", "for fire"}
	x.call()
	x.show()

	h := huawei{"hw p40", "your better phone"}
	h.call()
	h.show()
}
