package lib

import "fmt"

// Run an Orion Abstract Syntax Tree (AST) generated by Parse
func Run(ast *Orion) error {
	functions := make(map[string]func(args []string))
	functions["print"] = func(args []string) {
		for i := 0; i < len(args); i++ {
			fmt.Print(args[i], " ")
		}
		fmt.Println("\b")
	}
	functions["input"] = func(args []string) {}

	//for i := 0; i < len(ast.FuncDefs); i++ {
	//	funcDef := ast.FuncDefs[i]
	//
	//	functions[funcDef.Name] = funcDef.Body
	//}≠

	for i := 0; i < len(ast.FuncCalls); i++ {
		funcCall := ast.FuncCalls[i]

		functions[funcCall.Name](funcCall.Params)
	}

	return nil
}
