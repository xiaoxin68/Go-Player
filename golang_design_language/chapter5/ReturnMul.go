package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _,url := range os.Args[1:]{
		links ,err := findLinks(url)
		if err!=nil {
			fmt.Fprintf(os.Stderr,"findLinks:%v\n",err)
			continue
		}
		for _,link := range links{
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
	resp,err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		return nil,err
	}
	if resp.StatusCode!=http.StatusOK {
		return nil,fmt.Errorf("getting %s:%s",url,resp.StatusCode)
	}
	doc,err:=html.Parse(resp.Body)
	if err!=nil {
		return nil,fmt.Errorf("parseing %s as html:%v",url,err)
	}
	return visit(nil,doc),nil
}

func visit(t interface{}, doc interface{}) []string {
	
}
