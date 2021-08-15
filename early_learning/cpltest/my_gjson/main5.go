package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func main() {
	gjson.AddModifier("myTo", func(json, arg string) string {
		if arg == "title" {
			return strings.ToTitle(json)
		}

		return json
	})

	const json = `{"children": ["hello", "world", "xiaomotong"]}`
	fmt.Println(gjson.Get(json, "children|@myTo:title"))
}
