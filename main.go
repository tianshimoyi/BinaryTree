package main

import (
	binary "DataStructure/BinaryTree/BinaryTreeStruct"
	"fmt"
)

func test() {
	root := binary.CreateComBT([]interface{}{"A", "B", "C", "D", "E", "F"})
	fmt.Println(root == nil)
	fmt.Println(root.Data)
	fmt.Println(root.LeftChild.Data)
	//var a []interface{}
	root.PreorderTraversal()
	fmt.Printf("前序遍历：%#v\n", binary.OrderResults)
	binary.OrderResults = []interface{}{}
	root.MiddleTraversal()
	fmt.Printf("中序遍历：%#v\n", binary.OrderResults)
	binary.OrderResults = []interface{}{}
	root.PostTraversal()
	fmt.Printf("后续遍历：%#v\n", binary.OrderResults)
	fmt.Println(root.GetDeep())
	n := 0
	root.GetLeafNode(&n)
	fmt.Println(n)
	root.ReverseBT()
	binary.OrderResults = []interface{}{}
	root.PreorderTraversal()
	fmt.Printf("前序遍历：%#v\n", binary.OrderResults)
	n = 0
	root.GetChildNode(&n)
	fmt.Println(n)
	fmt.Println(root.IsFullBT())
}

func main() {
	test()
}
