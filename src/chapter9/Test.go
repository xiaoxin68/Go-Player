package main

import "fmt"

func main() {
	users := make(map[string]map[string]string)
	users["lily"] = make(map[string]string, 2)
	users["lily"]["password"] = "qwert"
	users["lily"]["nickname"] = "大佬"
	users["jack"] = map[string]string{
		"password": "rewsd",
		"nickname": "小仙女",
	}
	fmt.Println(users)
	modifyUser(users, "tom")
	fmt.Println(users)
	modifyUser(users, "tom")
	fmt.Println(users)
}
func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["password"] = "888888"
	} else {
		users[name] = map[string]string{
			"password": "123456",
			"nickname": "小菜鸡",
		}
	}
}
