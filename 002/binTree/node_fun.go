package binTree

import (
	"log"
	"regexp"
	"strconv"
)

func (this *NodeStar) SetObj(Obj any) *NodeStar {
	this.Obj = Obj
	return this
}

func (this *NodeStar) SetParent(parent *NodeStar) *NodeStar {
	this.Parent = parent
	return this
}

func (this *NodeStar) SetLeftChild(Left_child *NodeStar) *NodeStar {
	this.Left_child = Left_child
	return this
}

func (this *NodeStar) SetRightChild(Right_child *NodeStar) *NodeStar {
	this.Right_child = Right_child
	return this
}

func (this *NodeStar) CheckAndBuild() *NodeStar {
	if !this.isItValidKey() {
		log.Println("[Fail][Build][NodeStar]", this)
		return nil
	}
	log.Println("[Build][NodeStar]", this)
	return this
}

func (this *NodeStar) isItValidKey() bool {
	re, err := regexp.Compile(`^\d+$`)
	if err != nil {
		log.Println("[Fail][Check][NodeStar][Key] " + err.Error())
	}
	return re.MatchString(strconv.Itoa(this.Key))

}
