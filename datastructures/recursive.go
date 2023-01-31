package main

import "fmt"
/*
在计算机科学中，分治法是一种很重要的算法。
字面上的解释是分而治之，就是把一个复杂的问题分成两个或更多的相同或相似的子问题。
直到最后子问题可以简单的直接求解，原问题的解即子问题的解的合并。
分治法一般使用递归来求问题的解。
*/

func Rescuvie(n int) int { ///求阶乘
	if n == 0 {
		return 1
	}
	return n * Rescuvie(n-1)
}

func main() {
	fmt.Println(Rescuvie(5))
}

/*
Rescuvie(5)
{5 * Rescuvie(4)}
{5 * {4 * Rescuvie(3)}}
{5 * {4 * {3 * Rescuvie(2)}}}
{5 * {4 * {3 * {2 * Rescuvie(1)}}}}
{5 * {4 * {3 * {2 * 1}}}}
{5 * {4 * {3 * 2}}}
{5 * {4 * 6}}
{5 * 24}
120
*/
////////////////////////////////////////////////////////////////
//222222 尾部递归是指递归函数在调用自身后直接传回其值，而不对其再加运算，效率将会极大的提高。
func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}
	return RescuvieTail(n-1, a*n)
}

func main() {
	fmt.Println(RescuvieTail(5, 1))
}

/*
RescuvieTail(5, 1)
RescuvieTail(4, 1*5)=RescuvieTail(4, 5)
RescuvieTail(3, 5*4)=RescuvieTail(3, 20)
RescuvieTail(2, 20*3)=RescuvieTail(2, 60)
RescuvieTail(1, 60*2)=RescuvieTail(1, 120)

*/

//33333333333333类似的累加和
function tailrecsum(x, running_total=0) {
    if (x===0) {
        return running_total;
    } else {
        return tailrecsum(x-1, running_total+x);
    }
}

//444444444444斐波那契数列

func F(n int, a1, a2 int) int {
    if n == 0 {
        return a1
    }
    return F(n-1, a2, a1+a2)
}

func main() {
    fmt.Println(F(1, 1, 1))
    fmt.Println(F(2, 1, 1))
    fmt.Println(F(3, 1, 1))
    fmt.Println(F(4, 1, 1))
    fmt.Println(F(5, 1, 1))
}

/////55555555二分查找
// 二分查找递归解法
func BinarySearch(array []int, target int, l, r int) int {
    if l > r {
        // 出界了，找不到
        return -1
    }

    // 从中间开始找
    mid := (l + r) / 2
    middleNum := array[mid]

    if middleNum == target {
        return mid // 找到了
    } else if middleNum > target {
        // 中间的数比目标还大，从左边找
        return BinarySearch(array, target, l, mid-1)
    } else {
        // 中间的数比目标还小，从右边找
        return BinarySearch(array, target, mid+1, r)
    }

}

func main() {
    array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
    target := 500
    result := BinarySearch(array, target, 0, len(array)-1)
    fmt.Println(target, result)

    target = 189
    result = BinarySearch(array, target, 0, len(array)-1)
    fmt.Println(target, result)
}



//////666666二分查找非递归
// 二分查找非递归解法
func BinarySearch2(array []int, target int, l, r int) int {
    ltemp := l
    rtemp := r

    for {
        if ltemp > rtemp {
            // 出界了，找不到
            return -1
        }

        // 从中间开始找
        mid := (ltemp + rtemp) / 2
        middleNum := array[mid]

        if middleNum == target {
            return mid // 找到了
        } else if middleNum > target {
            // 中间的数比目标还大，从左边找
            rtemp = mid - 1
        } else {
            // 中间的数比目标还小，从右边找
            ltemp = mid + 1
        }
    }
}

func main() {
    array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
    target := 500
    result := BinarySearch2(array, target, 0, len(array)-1)
    fmt.Println(target, result)

    target = 189
    result = BinarySearch2(array, target, 0, len(array)-1)
    fmt.Println(target, result)
}


