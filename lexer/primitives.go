package lexer

import "unicode"

type Peekable interface {
	PeekRunes(uint) ([]rune, error)
	PeekSingleRune() (rune, error)
}

type Readable interface {
	ReadRunes(uint) ([]rune, error)
	ReadTill(func(rune) bool) []rune
}

type CommentLexer struct{}

func NewCommentLexer() *CommentLexer {
	return &CommentLexer{}
}

// Match matches if 1st and 2nd characters are /, ie //.
func (c *CommentLexer) Match(p Peekable) bool {
	chars, err := p.PeekRunes(2)
	if err != nil {
		return false
	}

	return string(chars) == "//"
}

// Lex lexes from after '//' to end of line. It does NOT implement multi line
// comments.
//
// TODO: implement multi line comments, ie // line1\n // line2.
func (c *CommentLexer) Lex(r Readable) []rune {
	r.ReadRunes(2) // throwaway '//' at beginning of line

	return r.ReadTill(
		func(char rune) bool { return char != '\n' },
	)
}

type NumberLexer struct{}

func NewIntegerLexer() *NumberLexer {
	return &NumberLexer{}
}

// Match matches if first character is a number.
func (i *NumberLexer) Match(p Peekable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	return unicode.IsNumber(char)
}

// Lex lexes integers and floats.
func (i *NumberLexer) Lex(r Readable) []rune {
	hasDot := false

	return r.ReadTill(
		func(char rune) bool {
			if unicode.IsNumber(char) {
				return true
			}

			if char == '.' {
				if hasDot {
					return false // already has a dot, so this is a method call
				} else {
					hasDot = true
					return true
				}
			}

			return false
		},
	)
}

type StringLexer struct{}

func NewStringLexer() *StringLexer {
	return &StringLexer{}
}

// Match matches if first character is double quotes.
func (s *StringLexer) Match(p Peekable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	return char == '"'
}

// Lex lexes all characters inside double quotes. It works with multiple line
// string, but NOT if any " are escaped	inside the string.
//
// TODO: implement escaping " inside double quotes, ie "Hello \" World"
func (s *StringLexer) Lex(r Readable) []rune {
	r.ReadRunes(1) // throwaway " at beginning of line

	chars := r.ReadTill(
		func(char rune) bool { return char != '"' },
	)

	r.ReadRunes(1) // throwaway " at end of line

	return chars
}
