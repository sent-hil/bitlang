package lexer

import (
	"bufio"
	"errors"
	"io"
)

var (
	ErrBoundsExceeded = errors.New("lexer: size is greater than slice length")
	ErrNegativeCount  = errors.New("lexer: negative count")
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

func (l *Lexer) PeekRunes(size int) ([]rune, error) {
	if size < 0 {
		return nil, ErrNegativeCount
	}
	if size > len(l.runes) {
		return nil, ErrBoundsExceeded
	}

	return l.runes[:size], nil
}

func (l *Lexer) ReadRunes(size int) ([]rune, error) {
	if size < 0 {
		return nil, ErrNegativeCount
	}
	if size > len(l.runes) {
		return nil, ErrBoundsExceeded
	}

	runes := l.runes[0:size]
	l.runes = l.runes[size:]

	return runes, nil
}
