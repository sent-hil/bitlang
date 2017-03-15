package lexer

import "unicode"

type Peekable interface {
	PeekRunes(uint) ([]rune, error)
	PeekSingleRune() (rune, error)
}

type Readable interface {
	ReadTill(func(rune) bool) []rune
}

type CommentLexer struct{}

func NewCommentLexer() *CommentLexer {
	return &CommentLexer{}
}

func (c *CommentLexer) Match(p Peekable) bool {
	runes, err := p.PeekRunes(2)
	if err != nil {
		return false
	}

	return string(runes) == "//"
}

func (c *CommentLexer) Lex(r Readable) (commentRunes []rune) {
	return r.ReadTill(
		func(char rune) bool { return char != '\n' },
	)
}

type IntegerLexer struct{}

func NewIntegerLexer() *IntegerLexer {
	return &IntegerLexer{}
}

func (i *IntegerLexer) Match(p Peekable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	return unicode.IsNumber(char)
}

func (i *IntegerLexer) Lex(r Readable) []rune {
	return r.ReadTill(
		func(char rune) bool { return unicode.IsNumber(char) },
	)
}
