package lib

type (
	// Orion struct.
	//
	// Define an Orion AST.
	Orion struct {
		FuncDefs  []*FuncDef  `@@*`
		FuncCalls []*FuncCall `@@*`
	}

	// FuncDef struct.
	//
	// Define a virtual function.
	FuncDef struct {
		Fn     string   `"fn"`
		Name   string   `@Ident`
		Params []string `"[" @Ident "]"`
		Body   *Body    `"{" @@ "}"`
	}

	// Body of a function or block.
	Body struct {
		Statements []*FuncCall `@@*`
	}

	// FuncCall struct.
	//
	// Call a function.
	FuncCall struct {
		Name   string   `@Ident`
		Params []string `"(" @String ")"`
	}
)
