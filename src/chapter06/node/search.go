package main

import (
	"errors"
	"fmt"
)

type BinarySearchTree struct {
	root *Node
}

// NewBinarySearchTree 初始化二叉排序树
func NewBinarySearchTree(node *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: node,
	}
}

// Insert 插入数据到二叉排序树
func (tree *BinarySearchTree) Insert(data interface{}) error {
	var value int
	var ok bool
	if value, ok = data.(int); !ok {
		return errors.New("暂时只支持int类型数据")
	}
	if node := tree.root; node == nil {
		// 二叉排序树为空，则初始化
		tree.root = NewNode(value)
		return nil
	} else {
		// 否则循环找到要插入的位置插入新的节点
		for node != nil {
			if value < node.Data.(int) {
				if node.Left == nil {
					node.Left = NewNode(value)
					break
				}
				node = node.Left
			} else if value > node.Data.(int) {
				if node.Right == nil {
					node.Right = NewNode(value)
					break
				}
				node = node.Right
			} else {
				return errors.New("对应数据已存在，请勿重复插入")
			}
		}
		return nil
	}
}

// Find 查找值为data的节点
func (tree *BinarySearchTree) Find(data int) *Node {
	if node := tree.root; node == nil {
		// 二叉排序树为空返回空
		return nil
	} else {
		// 否则返回值等于data的节点指针
		for node != nil {
			if data < node.Data.(int) {
				node = node.Left
			} else if data > node.Data.(int) {
				node = node.Right
			} else {
				return node
			}
		}
		return nil
	}
}

// Delete 删除值为 data 的节点
func (tree *BinarySearchTree) Delete(data int) error {
	if tree.root == nil {
		return errors.New("二叉树为空")
	}

	p := tree.root
	var pp *Node = nil // p 的父节点

	// 找到待删除节点
	for p != nil && p.Data.(int) != data {
		pp = p
		if p.Data.(int) < data {
			// 当前节点值小于待删除值，去右子树查找
			p = p.Right
		} else {
			// 否则去左子树查找
			p = p.Left
		}
	}
	if p == nil {
		return errors.New("待删除节点不存在")
	}

	// 待删除节点有两个子节点，需要将待删除节点值设置为该节点右子树最小节点值，再删除右子树最小节点
	if p.Left != nil && p.Right != nil {
		minP := p.Right // 右子树中的最小节点
		minPP := p      // minP 的父节点
		// 查找右子树中的最小节点
		for minP.Left != nil {
			minPP = minP
			minP = minP.Left
		}
		p.Data = minP.Data // 将 minP 的数据设置到 p 中
		p = minP           // 下面就变成删除 minP 了
		pp = minPP
	}

	// 应用待删除节点只有左/右子节点的逻辑
	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	if pp == nil {
		// 删除跟节点
		tree.root = nil
	} else if pp.Left == p {
		// 仅有左子节点
		pp.Left = child
	} else if pp.Right == p {
		// 仅有右子节点
		pp.Right = child
	}

	return nil
}

func main() {
	tree := NewBinarySearchTree(nil)
	_ = tree.Insert(3)
	_ = tree.Insert(2)
	_ = tree.Insert(5)
	_ = tree.Insert(1)
	_ = tree.Insert(4)

	fmt.Println("中序遍历二叉排序树：")
	midOrderTraverse(tree.root)
	fmt.Println()

	fmt.Println("查找值为 3 的节点：")
	node := tree.Find(3)
	fmt.Printf("%v\n", node)
	fmt.Println("查找值为 10 的节点：")
	node = tree.Find(10)
	fmt.Printf("%v\n", node)

	fmt.Println("删除值为 3 的节点：")
	err := tree.Delete(3)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("删除成功")
	}

	fmt.Println("中序遍历删除节点后的二叉排序树：")
	midOrderTraverse(tree.root)
	fmt.Println()
}
