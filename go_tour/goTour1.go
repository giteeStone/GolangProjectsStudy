package main

import (
	"fmt"
	"math/rand" //D:\Program Files\Go\src\math\rand\rand.go
	//包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始
)

func main() {
	fmt.Printf("My favorite number is %d", rand.Intn(20)) //随机数与机器有关
}
