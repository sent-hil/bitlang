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
	chars, err := p.PeekRunes(2)
	if err != nil {
		return false
	}

	return string(chars) == "//"
}

func (c *CommentLexer) Lex(r Readable) []rune {
	return r.ReadTill(
		func(char rune) bool { return char != '\n' },
	)
}

type NumberLexer struct{}

func NewIntegerLexer() *NumberLexer {
	return &NumberLexer{}
}

func (i *NumberLexer) Match(p Peekable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	return unicode.IsNumber(char)
}

func (i *NumberLexer) Lex(r Readable) []rune {
	hasDot := false

	return r.ReadTill(
		func(char rune) bool {
			if unicode.IsNumber(char) {
				return true
			}

			if char == '.' {
				if hasDot {
					return false // already has a dot, so this is an error
				} else {
					hasDot = true
					return true
				}
			}

			return false
		},
	)
}
