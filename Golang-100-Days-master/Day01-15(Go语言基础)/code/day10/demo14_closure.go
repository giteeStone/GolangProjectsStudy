package main

import "fmt"

func main() {
	/*
		go语言支持函数式编程：
			支持将一个函数作为另一个函数的参数，
			也支持将一个函数作为另一个函数的返回值。

		闭包(closure)：
			一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量(外层函数中的参数，或者外层函数中直接定义的变量)，并且该外层函数的返回值就是这个内层函数。

			这个内层函数和外层函数的局部变量，统称为闭包结构。


			局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
			但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。


	*/
	res1 := increment()        //res1 = fun,捕获外层变量i
	fmt.Printf("%T\n", res1)   //func() int
	fmt.Println("res1:", res1) //res1: 0x1070280
	v1 := res1()
	fmt.Println("v1:", v1) //1
	v2 := res1()
	fmt.Println("v2:", v2) //2
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	res2 := increment()
	fmt.Println("res2:", res2) //res2: 0x1070260 res2 = fun,另外一块fun函数块，捕获另一个外层变量i
	v3 := res2()
	fmt.Println("v3:", v3) //1
	fmt.Println(res2())

	fmt.Println(res1())
}

func increment() func() int { //外层函数
	//1.定义了一个局部变量
	i := 0
	//2.定义了一个匿名函数，给变量自增并返回
	fun := func() int { //内层函数
		i++
		return i
	}
	//fmt.Print("fun:", fun) //func() int
	//3.返回该匿名函数
	return fun
}
