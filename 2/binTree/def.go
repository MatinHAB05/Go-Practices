package binTree

type NodeStar struct {
	Obj    any
	Key    int
	Parent *NodeStar
	// children []NodeStar
	Left_child  *NodeStar
	Right_child *NodeStar
}

type Tree struct {
	Head *NodeStar
	key  int
}
