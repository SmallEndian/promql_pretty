package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/prometheus/prometheus/promql/parser"
)

func main() {

	var indent_level int

	flag.IntVar(&indent_level, "indent_level", 0, "Level information to determine at which level/depth the current node is in the AST")
	flag.IntVar(&indent_level, "i", 0, "Level information to determine at which level/depth the current node is in the AST")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	var input strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input.WriteString(line)

	}

	expr, err := parser.ParseExpr(input.String())

	if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	fmt.Print(expr.Pretty(indent_level))

}
