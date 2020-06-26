package biotree

import(
	 "fmt"
)

type Tree struct{
	
	Root *Node
}


func (t *Tree) AddBranch(id string)*Node{

	node := &Node{}

	if node == nil{
	//Root case
		if node.RightBranch == nil && node.LeftBranch == nil{
			node = node.ChildNode(id)
			fmt.Println("Root added:")
		}else{
			// if node.LeftBranch == nil{
			// 	lChild := node.ChildNode(id)
			// 	lChild.LeftBranch = node
			// 	lChild.RightBranch = nil
			// 	fmt.Println("LeftBranch added:")	
			// }
		}
	}
	return node
}