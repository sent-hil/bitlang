package lexer

import "github.com/sent-hil/bitlang/runeio"

type Peekable interface {
	PeekRunes(uint) ([]rune, error)
	PeekSingleRune() (rune, error)
}

type Readable interface {
	Peekable
	ReadRunes(uint) ([]rune, error)
	ReadTill(func(rune) bool) []rune
}

type Lexer struct {
	reader *runeio.Reader
}

func NewLexer(reader *runeio.Reader) (*Lexer, error) {
	return &Lexer{reader: reader}, nil
}
