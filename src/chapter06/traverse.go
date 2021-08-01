package main

import "fmt"

// Node 通过链表存储二叉树节点信息
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}
func (node *Node) GetData() string {
	return fmt.Sprintf("%v", node.Data)
}

// 前序遍历
func preOrderTraverse(treeNode *Node) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先打印当前节点值
	fmt.Printf("%s ", treeNode.GetData())
	// 然后对左子树和右子树做前序遍历
	preOrderTraverse(treeNode.Left)
	preOrderTraverse(treeNode.Right)
}

// 测试代码
func main() {
	// 初始化一个简单的二叉树
	node1 := NewNode(0) // 根节点
	node2 := NewNode("1")
	node3 := NewNode(2.0)
	node1.Left = node2
	node1.Right = node3
	// 前序遍历这个二叉树
	fmt.Print("前序遍历: ")
	preOrderTraverse(node1)
	fmt.Println()
}
