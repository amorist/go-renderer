package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type User struct {
	Name string
}

func main() {
	tpl := template.New("")
	tpl, _ = tpl.Parse("<p> hello {{.Name}} </p>")
	data := User{Name: "Tom"}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String()) // 渲染后的字符串 // <p> hello Tom </p>
}
