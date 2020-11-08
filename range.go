package main

/*
    author:qingbing
    file:range
    function:study range for golang
*/
import "fmt"


func main(){

//1、range 对于 数组、切片

    s := []string{"apple","pen"}
    for i,value := range s{
        fmt.Println(i,value)
    }

//2、对于字符串
    for i,value := range "hello"{
        fmt.Println(i,value)
    }

//3、range对于map集合
    m := map[string]string{"name":"xiaopang","age":"25"}
    for i,value := range m{
        fmt.Println(i,value)
    }

//4、占位符_
    sum := 0
    nums := []int{1,2,3,4,5}
    for _,value := range nums{
        sum += value
    }
    fmt.Println(sum)
}
