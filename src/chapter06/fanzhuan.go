package main

import "fmt"

type Node1 struct {
	data interface{}
	next *Node1
}

// 反转单链表
func (head *Node1) reverse() *Node1 {
	//  空链表
	if head == nil {
		return nil
	}

	var reverseHead *Node1 // 反转后单链表的头结点

	var preNode *Node1
	curNode := head

	for curNode != nil {
		nextNode := curNode.next
		if nextNode == nil {
			reverseHead = curNode // 尾结点转换为头结点
		}
		// 反转实现，当前结点的前驱结点变成后驱结点
		curNode.next = preNode
		// 设置下一个结点的前驱结点
		preNode = curNode
		curNode = nextNode
	}
	// 返回反转后的头结点
	return reverseHead
}

// 遍历单链表
func (head *Node1) traverse() {
	node := head
	for node != nil {
		fmt.Printf("%v ", node.data)
		node = node.next
	}
}

/* func main() {
	first := &Node1{data: 1}
	second := &Node1{data: 2}
	third := &Node1{data: 3}

	first.next = second
	second.next = third

	head := first

	fmt.Print("反转前: ")
	head.traverse()
	fmt.Println()

	newHead := head.reverse()

	fmt.Print("反转后: ")
	newHead.traverse()
	fmt.Println()
} */
