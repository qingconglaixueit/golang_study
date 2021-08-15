/*
	author:qingbing
	date: 2020年11月 8日 21:12:01
	what: json数据 -> go数据对象   //对于已知json结构进行解析
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//1.根据已知json文件进行分析，定义struct

type Post struct {
	Id        int         `json:"id"`
	Content   string      `json:"content"`
	Author    *Author     `json:"author"`
	Published bool        `json:"published"`
	Label     []string    `json:"label"`
	NextPost  *Post       `json:"nextPost"`
	Comments  []*Comments `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comments struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {

	//2.使用os 打开json数据文件
	fh, err := os.Open("./data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()

	//3.读取json文件，将数据保存下来
	jsondata, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(jsondata))
	//4.解析jsondata 将数据放到 结构体Post中
	var post Post
	err = json.Unmarshal(jsondata, &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	//5.使用解析出来的数据
	fmt.Println(post)
	fmt.Println(post.Author.Id)
	fmt.Println(post.Author.Name)

	fmt.Println(post.Comments[0].Id)
	fmt.Println(post.Comments[0].Author)
	fmt.Println(post.Comments[0].Content)
}
