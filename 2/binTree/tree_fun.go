package binTree

import "log"

func InitTree() *Tree {
	t := &Tree{}
	return t
}
func (this *Tree) SetHead(Head *NodeStar) *Tree {
	this.Head = Head
	return this
}

func (this *Tree) Build() *Tree {
	log.Println("[Build][Tree]", this)
	return this

}

func (this *Tree) makeNewKey() int {
	k := this.key
	this.key++
	return k
}

func (this *Tree) InitNode() *NodeStar {
	node := &NodeStar{}
	node.Key = this.makeNewKey()
	return node
}
