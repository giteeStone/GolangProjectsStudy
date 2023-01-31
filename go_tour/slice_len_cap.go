package main

import "fmt"

func main() {
	s1 := []int{2, 3, 5, 7, 11, 13}
	s := s1
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s1)
	s = append(s, 111)
	s = append(s, 222)
	printSlice(s)

	printSlice(s1)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
