package main

import (
	"fmt"
)

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	// 新的节点
	node := new(LinkNode)
	node.Data = 2

	// 新的节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1 // node1 链接到 node 节点上

	// 新的节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2 // node2 链接到 node1 节点上

	// 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印节点值
			fmt.Println(nowNode.Data)
			// 获取下一个节点
			nowNode = nowNode.NextNode
		} else {
			break
		}

	}
}

// 循环链表
type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func main() {
	r := new(Ring)
	r.init()
}
