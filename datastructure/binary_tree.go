package datastructure

import(
	"fmt"
)

//Node
type Node struct {
	Value int
	Left *Node
	Right *Node
}

func main(){
	/*
					8
			1				7
		3		5		6
	*/

	node1 := &Node{Value:6}
	node2 := &Node{Value: 7, Left: node1}

	node3 := &Node{Value: 5}
	node4 := &Node{Value: 3}
	node5 := &Node{Value: 1, Right:node3, Left:node4}
	node6 := &Node{Value: 8, Right: node2, Left: node5}

	fmt.Println(leaves(node6))

}

func leaves(tree *Node) []int {
	result := make([]int, 0)
	if tree == nil {
		return result
	}
	
	if tree.Left == nil && tree.Right == nil {
		return append(result, tree.Value)
	}
	
	result = append(result, leaves(tree.Left)...)
	result = append(result, leaves(tree.Right)...)

	return result
}