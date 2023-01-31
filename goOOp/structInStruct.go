package main

import "fmt"

type Goods struct {
	Name  string
	Price float32
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

// 也可以使用这种地址方式的匿名结构体
type Computer struct {
	*Goods
	*Brand
}

func main() {

	tv := TV{
		Goods{
			Name:  "电视机",
			Price: 10000,
		},
		Brand{
			Name:    "海尔",
			Address: "全世界",
		},
	}

	fmt.Println("tv:", tv)

	computer := Computer{
		&Goods{
			Name:  "计算机",
			Price: 10000,
		},
		&Brand{
			Name:    "海尔",
			Address: "全世界",
		},
	}

	fmt.Println("computer:", computer)
	fmt.Println(*(computer.Goods), *(computer.Brand))
}
