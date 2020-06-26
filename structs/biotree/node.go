package biotree




type Node struct{
	Object string
	RightBranch *Node
	LeftBranch  *Node
}


func (n *Node) ChildNode(id string) *Node{

	if id != ""{
		(*n).Object = id
		(*n).RightBranch = nil	
		(*n).LeftBranch = nil
	}

	return n
}