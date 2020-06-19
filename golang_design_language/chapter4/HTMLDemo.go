package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>A:{{.A}}</p><p>B:{{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct{
		A string //不受信任的纯文本
		B template.HTML  //受信任的html
	}
	data.A = "<b>hello</b>"
	data.B = "<b>hello</b>"
	if err:=t.Execute(os.Stdout,data);err != nil {
		log.Fatal(err)
	}

}
//go run HTMLDemo.go repo:golang/go comment:gopherbot json encoder >issues.html
//file:///D:/development/jetbrains/goland/workspace/src/Go-Player/golang_design_language/chapter4/issues.html
