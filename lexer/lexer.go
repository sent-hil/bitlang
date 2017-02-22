package lexer

import (
	"errors"

	"github.com/sent-hil/bitlang/runeio"
)

var (
	ErrBoundsExceeded = errors.New("lexer: size is greater than slice length")
	ErrNegativeCount  = errors.New("lexer: negative count")
)

type Lexer struct {
	reader *runeio.Reader
}

func NewLexer(reader *runeio.Reader) (*Lexer, error) {
	return &Lexer{reader: reader}, nil
}
