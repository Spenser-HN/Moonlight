package syntax

import "fmt"

func CreateAST(SourceFiles TokenizedFiles) {

	var ResultAst *AST = &AST{
		Type: "Program",
		Body: []Body{
			{
				Type:         "VariableDeclaration",
				Declarations: []string{},
			},
		},
	}

	fmt.Println(SourceFiles, "\n\n AST:\n", ResultAst)
}
