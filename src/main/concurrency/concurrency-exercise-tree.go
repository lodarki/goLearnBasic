package main

import "fmt"

//可以用多种不同的二叉树的叶子节点存储相同的数列值。例如，这里有两个二叉树保存了序列 1，1，2，3，5，8，13。
//
//
//用于检查两个二叉树是否存储了相同的序列的函数在多数语言中都是相当复杂的。这里将使用 Go 的并发和 channel 来编写一个简单的解法。
//
//这个例子使用了 tree 包，定义了类型：

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func addTree(handleValue int, tree *Tree) {

	if handleValue <= tree.Value {
		if tree.Left != nil {
			addTree(handleValue, tree.Left)
		} else {
			tree.Left = &Tree{nil, handleValue, nil}
		}
	}
	if handleValue > tree.Value {
		if tree.Right != nil {
			addTree(handleValue, tree.Right)
		} else {
			tree.Right = &Tree{nil, handleValue, nil}
		}
	}
	//fmt.Println(tree)
}

func buildTrees(vs []int) *Tree {

	if len(vs) == 0 {
		return nil
	}

	rootValue := vs[0]
	rootTree := &Tree{nil, rootValue, nil}
	if len(vs) > 1 {
		for _, v := range vs[1:] {
			addTree(v, rootTree)
		}
	}
	return rootTree
}

func collectTreeValues(tree *Tree) []int {
	vsChan := make(chan int, 100)
	go collectLeafValue(tree, vsChan)
	var vs []int
	//for element := range vsChan {
	//	vs = append(vs, element)
	//}
	return vs
}

func collectLeafValue(tree *Tree, vschan chan int) {
	vschan <- tree.Value
	fmt.Println(len(vschan))
	if tree.Left != nil {
		go collectLeafValue(tree.Left, vschan)
	}
	if tree.Right != nil {
		go collectLeafValue(tree.Right, vschan)
	}
}

func main() {
	nodes1 := []int{3, 1, 8, 5, 13, 1, 2}
	nodes2 := []int{8, 13, 3, 1, 1, 2, 5}
	trees1 := buildTrees(nodes1)
	trees2 := buildTrees(nodes2)
	fmt.Println(trees1) //树构建成功
	fmt.Println(trees2) //树构建成功

	fmt.Println(collectTreeValues(trees1))
	fmt.Println(collectTreeValues(trees2))
}
