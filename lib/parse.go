package lib

import (
	"fmt"
	"os"

	"github.com/alecthomas/participle/v2"
)

// Parse a file to Orion AST.
func Parse(file string) (*Orion, error) {
	// Build the parser
	parser, err := participle.Build[Orion](participle.Unquote())
	if err != nil {
		return nil, fmt.Errorf("failed to build parser")
	}

	// Read the file
	contents, err := os.ReadFile("examples/main.or")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s", file)
	}

	// Parse an example program
	ast, err := parser.ParseBytes("main.or", contents)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %s", file)
	}

	return ast, nil
}
