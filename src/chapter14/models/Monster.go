package models

//对于结构体的序列化，如果希望序列化后的名字由我们重新制定，可以给一个tag标签
type Monster struct {
	Name     string  `json:"monster_name"`
	Age      int `json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}
