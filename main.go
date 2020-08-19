package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/skanehira/dentaku/lexer"
	"github.com/skanehira/dentaku/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		print(">> ")
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseError(out, p.Errors())
			continue
		}

		fmt.Fprintln(out, program.String())
	}
}

func printParseError(out io.Writer, errors []string) {
	for _, e := range errors {
		fmt.Fprintln(out, e)
	}
}

func main() {
	Start(os.Stdin, os.Stdout)
}
