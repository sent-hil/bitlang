package lexer

import "github.com/sent-hil/bitlang/runeio"

type Lexer struct {
	reader *runeio.Reader
}

func NewLexer(reader *runeio.Reader) (*Lexer, error) {
	return &Lexer{reader: reader}, nil
}
