package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// MatinHAB05
func main() {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(
		fset,
		"text.code.txt",
		nil,
		parser.ParseComments,
	)
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	fmt.Println("File parsed successfully")
	fmt.Println("Comments:", len(node.Comments))

	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return true
		}

		fmt.Printf("NODE: %T\n", n)

		if com, ok := n.(*ast.CommentGroup); ok {
			fmt.Println("COMMENT:", com.List)
		}

		return true
	})
}
