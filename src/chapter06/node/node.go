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

// 中序遍历
func midOrderTraverse(treeNode *Node) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先从左子树最左侧节点开始遍历
	midOrderTraverse(treeNode.Left)
	// 打印位于中间的根节点
	fmt.Printf("%s ", treeNode.GetData())
	// 最后按照和左子树一样的逻辑遍历右子树
	midOrderTraverse(treeNode.Right)
}
