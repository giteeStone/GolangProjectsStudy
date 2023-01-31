package main

import "fmt"

func main() {
	/*
		结构体嵌套：一个结构体中的字段，是另一个结构体类型。
			has a
	*/
	s2 := Student{name: "李小花", age: 19, book: Book{bookName: "Go语言是怎样炼成的", price: 89.7}}
	fmt.Println(s2.name, s2.age)
	fmt.Println("\t", s2.book.bookName, s2.book.price)

	s3 := Student{
		name: "Jerry",
		age:  17,
		book: Book{
			bookName: "十万个为啥",
			price:    55.9,
		},
	}
	fmt.Println(s3.name, s3.age)
	fmt.Println("\t", s3.book.bookName, s3.book.price)

	fmt.Println("------------->结构体值传递1，深拷贝1")
	b1 := Book{}
	b1.bookName = "西游记"
	b1.price = 45.8

	s1 := Student{}
	s1.name = "王二狗"
	s1.age = 18
	s1.book = b1 //值传递
	fmt.Println(b1)
	fmt.Println(s1)
	fmt.Printf("学生姓名：%s，学生年龄：%d，看的书是：《%s》，书的价格是:%.2f\n", s1.name, s1.age, s1.book.bookName, s1.book.price)

	fmt.Println("------------->结构体值传递2，深拷贝2")
	s1.book.bookName = "红楼梦"
	fmt.Println(s1)
	fmt.Println(b1)

	fmt.Println("------------->结构体引用传递，浅拷贝")
	b4 := Book{bookName: "呼啸山庄", price: 76.9}
	s4 := Student2{name: "Ruby", age: 18, book: &b4}
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println("\t", s4.book)
	fmt.Println("------------->结构体修改 被引用的嵌套结构体")
	s4.book.bookName = "雾都孤儿"
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println("\t", s4.book)

}

// 1.定义一个书的结构体
type Book struct {
	bookName string
	price    float64
}

// 2.定义学生的结构体
type Student struct {
	name string
	age  int
	book Book
}

type Student2 struct {
	name string
	age  int
	book *Book // book的地址(赋值时用引用传递，防止值传递的深拷贝)
}
