package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}
type Machine struct {
	*Goods
	*Brand
}

func main() {
	tv := TV{Goods{"goods1", 2345.89}, Brand{"brand1", "湖北武汉"}}
	fmt.Println(tv.Goods, tv.Brand)
	
	tv1 := TV{
		Goods: Goods{
			Name:  "goods2",
			Price: 1234.56,
		},
		Brand: Brand{
			Name:    "brand2",
			Address: "湖北随州",
		},
	}
	fmt.Println(tv1.Goods, tv1.Brand)
	
	machine := Machine{&Goods{"goods3", 1234.56}, &Brand{"brand3", "湖北仙桃"}}
	fmt.Println(*machine.Goods, *machine.Brand)
	
	machine2 := Machine{
		&Goods{
			Name:  "goods4",
			Price: 1234.56,
		},
		&Brand{
			Name:    "brand4",
			Address: "湖北荆门",
		},
	}
	fmt.Println(*machine2.Goods, *machine2.Brand)
}
