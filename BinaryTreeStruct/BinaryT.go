package BinaryTreeStruct

import (
	"fmt"
	"math"
	"reflect"
)

type BinaryTree struct {
	Data       interface{}
	LeftChild  *BinaryTree
	RightChild *BinaryTree
}

var (
	OrderResults []interface{}
)

//TODO 创建二叉树

func NewBinaryT(data interface{}) *BinaryTree {
	return &BinaryTree{
		Data: data,
	}
}

//TODO 添加左子树

func AddLeftChild(node *BinaryTree, data interface{}) {
	node.LeftChild = NewBinaryT(data)
}

//TODO 添加右子树

func AddRightChild(node *BinaryTree, data interface{}) {
	node.RightChild = NewBinaryT(data)
}

//TODO 前序遍历

func (root *BinaryTree) PreorderTraversal() {
	if root == nil {
		return
	} else {
		OrderResults = append(OrderResults, root.Data)
		fmt.Print(root.Data, " ")
		root.LeftChild.PreorderTraversal()
		root.RightChild.PreorderTraversal()
	}
	//fmt.Println(a)
}

//TODO 中序遍历

func (root *BinaryTree) MiddleTraversal() {
	if root == nil {
		return
	} else {
		root.LeftChild.MiddleTraversal()
		OrderResults = append(OrderResults, root.Data)
		root.RightChild.MiddleTraversal()
	}
}

//TODO 后续遍历

func (root *BinaryTree) PostTraversal() {
	if root == nil {
		return
	} else {
		root.LeftChild.PostTraversal()
		root.RightChild.PostTraversal()
		OrderResults = append(OrderResults, root.Data)
	}
}

//TODO 创建完全二叉树

func CreateComBT(data []interface{}) *BinaryTree {
	nodes := make([]*BinaryTree, len(data))
	for i := 0; i < len(data); i++ {
		nodes[i] = NewBinaryT(data[i])
	}
	for i := 0; i < len(data)/2; i++ {
		nodes[i].LeftChild = nodes[2*i+1]
		if 2*i+2 <= len(data)-1 {
			nodes[i].RightChild = nodes[2*i+2]
		}
	}
	return nodes[0]
}


//TODO查询二叉树深度

func (root *BinaryTree) GetDeep() int {
	if root == nil {
		return 0
	}
	lh := root.LeftChild.GetDeep()
	rh := root.RightChild.GetDeep()
	if lh > rh {
		lh++
		return lh
	} else {
		rh++
		return rh
	}
}

//TODO 统计叶节点

func (root *BinaryTree) GetLeafNode(n *int) {
	if root == nil {
		return
	}
	if root.LeftChild == nil && root.RightChild == nil {
		(*n)++
	}
	root.LeftChild.GetLeafNode(n)
	root.RightChild.GetLeafNode(n)
}

//TODO 统计子节点

func (root *BinaryTree) GetChildNode(n *int) {
	if root == nil {
		return
	}
	if root.LeftChild != nil || root.RightChild != nil {
		(*n)++
	}
	root.LeftChild.GetChildNode(n)
	root.RightChild.GetChildNode(n)
}

//TODO 判断二叉树是否为满二叉树

func (root *BinaryTree) IsFullBT() bool { //n=deep nodesum=2^n-1
	sum := 0
	n := 0
	root.GetChildNode(&n)
	sum += n
	n = 0
	root.GetLeafNode(&n)
	sum += n
	if sum == int(math.Pow(2, float64(root.GetDeep())))-1 {
		return true
	} else {
		return false
	}
}

//TODO 查找是否存在数据

func (root *BinaryTree) SearchData(data interface{}) {
	if root == nil {
		return
	}
	if reflect.DeepEqual(root.Data, data) {
		fmt.Println("find the data")
		return
	}
	root.LeftChild.SearchData(data)
	root.RightChild.SearchData(data)
}

//TODO 销毁二叉树

func (root *BinaryTree) DestroyBT() {
	if root == nil {
		return
	}
	root.LeftChild.DestroyBT()
	root.LeftChild = nil
	root.RightChild.DestroyBT()
	root.RightChild = nil
	root = nil
}

//TODO 二叉树翻转

func (root *BinaryTree) ReverseBT() {
	if root == nil {
		return
	}
	root.LeftChild, root.RightChild = root.RightChild, root.LeftChild
	root.LeftChild.ReverseBT()
	root.RightChild.ReverseBT()
}

//TODO 二叉树拷贝

func (root *BinaryTree) CopyBT() *BinaryTree {
	if root == nil {
		return nil
	}
	lChild := root.LeftChild.CopyBT()
	rChild := root.RightChild.CopyBT()
	node := &BinaryTree{
		Data:       root.Data,
		LeftChild:  lChild,
		RightChild: rChild,
	}
	return node
}
