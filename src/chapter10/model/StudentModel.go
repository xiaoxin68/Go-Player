package model

//如果model包中的Student首字母大写，别的包引入后直接使用没问题
type Student struct {
	Name string
	Age int
}

//如果model包中的Student首字母大写，别的包引入后不能直接使用，用工厂模式解决问题
type student struct {
	Name string
	Age int
}

func NewStudent(n string,a int)*student  {
	return &student{
		Name: n,
		Age:  a,
	}
}

//如果model包中的结构体中某字段首字母小写，别的包引入后如何处理呢：提供一个get方法
type student1 struct {
	Name string
	age int
}
func NewStudent1(n string,a int)*student1  {
	return &student1{
		Name: n,
		age:  a,
	}
}
func (s *student1) GetAge()int  {
	return s.age
}
func main() {
	
}
