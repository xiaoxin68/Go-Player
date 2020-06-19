package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}
func main() {
	var movies = []Movie{
		{Title: "asdsa",Year: 1993,Color: false,Actors: []string{"23r","s","adfw"}},
		{Title: "sad",Year: 1943,Color: true,Actors: []string{"asf","asda"}},
		{Title: "sd",Year: 1973,Color: false,Actors: []string{"qdad","axza","rdwewweqd"}},
	}
	//data, err := json.Marshal(movies)
	//分别是每行输出的前缀字符串；定义缩进的字符串
	data, err := json.MarshalIndent(movies,"","    ")
	if err!=nil {
		log.Fatal("JSON marshaling failed:%s",err)
	}
	fmt.Printf("%s\n",data)

	var res []Movie
	err = json.Unmarshal(data, &res)
	if err!=nil {
		log.Fatal("JSON ummarshaling failed:%s",err)
	}
	fmt.Println(res)
}
