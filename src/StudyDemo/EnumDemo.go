package main

import "fmt"

func main() {
	//Go语言现阶段没有枚举类型，但是可以使用 const 常量配合iota 来模拟枚举类型
	type Weapon int
	
	const (
		Arrow Weapon = iota // 开始生成枚举值, 默认为0
		/*1、将常量 Arrow 的类型标识为 Weapon，这样标识后，const 下方的常量可以使用 Weapon 作为默认类型。
		该行使用 iota 进行常量值自动生成，iota 的起始值为 0，
		一般情况下也是建议枚举从 0 开始，让每个枚举类型都有一个空值，方便业务和逻辑的灵活使用。
		2、一个 const 声明内的每一行常量声明，将会自动套用前面的 iota 格式，并自动增加，*/
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	
	// 输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower) //0 1 2 3 4
	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon) //4
	
	////////////////////////////////////////////////////////////
	
	//iota 不仅可以生成每次增加 1 的枚举值。还可以利用 iota 来做一些强大的枚举常量值生成器。下面的代码可以方便的生成标志位常量：
	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue) //2 4 8
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue) //10 100 1000  按照二进制输出
	
}
