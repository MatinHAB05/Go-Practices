package main

import (
	"test/binTree"
)

func main() {
	myTree := binTree.InitTree()

	// Level 4 (leaf nodes)
	n1 := myTree.InitNode().SetObj("Node-1").CheckAndBuild()
	n2 := myTree.InitNode().SetObj("Node-2").CheckAndBuild()
	n3 := myTree.InitNode().SetObj("Node-3").CheckAndBuild()
	n4 := myTree.InitNode().SetObj("Node-4").CheckAndBuild()
	n5 := myTree.InitNode().SetObj("Node-5").CheckAndBuild()
	n6 := myTree.InitNode().SetObj("Node-6").CheckAndBuild()
	n7 := myTree.InitNode().SetObj("Node-7").CheckAndBuild()
	n8 := myTree.InitNode().SetObj("Node-8").CheckAndBuild()
	n9 := myTree.InitNode().SetObj("Node-9").CheckAndBuild()
	n10 := myTree.InitNode().SetObj("Node-10").CheckAndBuild()
	n11 := myTree.InitNode().SetObj("Node-11").CheckAndBuild()
	n12 := myTree.InitNode().SetObj("Node-12").CheckAndBuild()
	n13 := myTree.InitNode().SetObj("Node-13").CheckAndBuild()
	n14 := myTree.InitNode().SetObj("Node-14").CheckAndBuild()
	n15 := myTree.InitNode().SetObj("Node-15").CheckAndBuild()

	// Level 3
	lA := myTree.InitNode().SetObj("L-A").SetLeftChild(n1).SetRightChild(n2).CheckAndBuild()
	lB := myTree.InitNode().SetObj("L-B").SetLeftChild(n3).SetRightChild(n4).CheckAndBuild()
	lC := myTree.InitNode().SetObj("L-C").SetLeftChild(n5).SetRightChild(n6).CheckAndBuild()
	lD := myTree.InitNode().SetObj("L-D").SetLeftChild(n7).SetRightChild(n8).CheckAndBuild()
	rA := myTree.InitNode().SetObj("R-A").SetLeftChild(n9).SetRightChild(n10).CheckAndBuild()
	rB := myTree.InitNode().SetObj("R-B").SetLeftChild(n11).SetRightChild(n12).CheckAndBuild()
	rC := myTree.InitNode().SetObj("R-C").SetLeftChild(n13).SetRightChild(n14).CheckAndBuild()
	rD := myTree.InitNode().SetObj("R-D").SetLeftChild(n15).SetRightChild(nil).CheckAndBuild()

	// Level 2
	leftHigh := myTree.InitNode().SetObj("Left-High").SetLeftChild(lA).SetRightChild(lB).CheckAndBuild()
	leftLow := myTree.InitNode().SetObj("Left-Low").SetLeftChild(lC).SetRightChild(lD).CheckAndBuild()
	rightHigh := myTree.InitNode().SetObj("Right-High").SetLeftChild(rA).SetRightChild(rB).CheckAndBuild()
	rightLow := myTree.InitNode().SetObj("Right-Low").SetLeftChild(rC).SetRightChild(rD).CheckAndBuild()

	// Level 1
	leftRoot := myTree.InitNode().SetObj("Left-Root").SetLeftChild(leftHigh).SetRightChild(leftLow).CheckAndBuild()
	rightRoot := myTree.InitNode().SetObj("Right-Root").SetLeftChild(rightHigh).SetRightChild(rightLow).CheckAndBuild()

	// Root
	head := myTree.InitNode().
		SetObj("Head").
		SetLeftChild(leftRoot).
		SetRightChild(rightRoot).
		CheckAndBuild()

	myTree.SetHead(head)
}
