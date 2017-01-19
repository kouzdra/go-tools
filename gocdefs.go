package main

import "os"
import "io"
import "strings"
import "fmt"
import "text/scanner"

const GoWhitespace = 1<<'\t' | /*1<<'\n' |*/ 1<<'\r' | 1<<' '
const GoTokens     = scanner.GoTokens & ^scanner.SkipComments
func main () {
	args := os.Args
	if args == nil || len (args) != 2 {
		fmt.Printf ("godefs: invalid argument #%d\n", len (args))
	} else {
		pfx := os.Args [1]
		fmt.Printf ("\n// %s*\n", pfx)
		scanDefs (os.Stdin, pfx)
	}
}

func scanDefs (reader io.Reader, pfx string) {
	var s scanner.Scanner
	s.Filename = "stdin"
	s.Init(reader)
	s.Whitespace = GoWhitespace
	var tok rune
	next := func (str string) bool {
		if tok != scanner.EOF && s.TokenText() == str {
			//fmt.Printf("OKK (%s) At position %s: [%s]\n", str, s.Pos(), s.TokenText())
			tok = s.Scan()
			return true
		}
		//fmt.Printf("NOT (%s) At position %s: [%s]\n", str, s.Pos(), s.TokenText())
		return false
	}
	for tok != scanner.EOF {
		if next ("\n") {
			if next ("#") && next ("define") {
				name := s.TokenText ()
				if strings.HasPrefix (name, pfx) {
					fmt.Printf ("const %s = C.%s\n", name, name)
				}
			}
		} else {
			if tok != scanner.EOF {
				tok = s.Scan ()
			}
		}
	}
}
