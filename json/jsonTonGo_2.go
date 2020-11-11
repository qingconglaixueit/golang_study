/*
	author:qingbing
	date: 2020年11月 8日 21:44:39
	what: json数据 -> go数据对象   //对于未知json结构进行解析
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//1.根据已知json文件进行分析，定义struct

func main() {

	//1.使用os 打开json数据文件
	fh, err := os.Open("./data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()

	//2.读取json文件，将数据保存下来
	jsondata, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(jsondata))
	//3.解析jsondata 将数据放到 结构体Post中
	var unknown interface{}
	err = json.Unmarshal(jsondata, &unknown)
	if err != nil {
		fmt.Println(err)
		return
	}
	//4.使用解析出来的数据  此时还是 map[string]interface{}
	//fmt.Println(unknown)

	//5.解析成可以控制的数据
	m := unknown.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "type:string", v)
			fmt.Println("--------------------")
		case float64:
			fmt.Println(k, "type:float64", v)
			fmt.Println("--------------------")
		case bool:
			fmt.Println(k, "type:bool", v)
			fmt.Println("--------------------")
		case map[string]interface{}:
			fmt.Println(k, "type:map[string]interface{}", v)
			for kk, value := range vv {
				fmt.Println(kk, ":", value)
			}
			fmt.Println("--------------------")
		case []interface{}:
			fmt.Println(k, "type:[]interface{}", v)
			for key, val := range vv {
				switch hh := val.(type) {
				case map[string]interface{}:
					fmt.Println(key, ":", val)
					for hk, hv := range hh {
						fmt.Println(hk, ":", hv)
					}
					fmt.Println("========")
				default:
					fmt.Println("========", hh)
				}

			}
			fmt.Println("--------------------")
		default:
			fmt.Println(k, "type:<nil>", v)
		}
	}

	//test := map[string]string{"name": "qingbing", "age": "20"}
	//fmt.Println(test)
	//fmt.Println(test["name"])
}
