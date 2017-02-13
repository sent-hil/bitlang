package lexer

import (
	"bufio"
	"io"
)

type Lexer struct {
	runes []rune
}

func NewLexer(reader io.Reader) (*Lexer, error) {
	runes := []rune{}
	bufReader := bufio.NewReader(reader)

	for {
		c, _, err := bufReader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		runes = append(runes, c)
	}

	return &Lexer{runes: runes}, nil
}
