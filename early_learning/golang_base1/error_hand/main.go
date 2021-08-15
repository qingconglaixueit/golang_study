package main

import "fmt"

//定义数据结构
type DivideError struct {
	devidee int
	devider int
}

//错误处理实现Error()接口
func (de *DivideError) Error() string {
	strdata := `
		error,divide is zero
		dividee is %d
		divider is zero
	`

	return fmt.Sprintf(strdata, de.devidee)
}

//实现功能接口
func Divide(dividee int, divider int) (result int, errMsg string) {
	if divider == 0 {
		data := DivideError{dividee, divider}
		errMsg = data.Error()
		return
	} else {
		return dividee / divider, ""
	}
}

func main() {
	a := 10
	b := 0
	result, err := Divide(a, b)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d / %d == %d \n", a, b, result)

}
