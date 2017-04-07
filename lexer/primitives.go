package lexer

import (
	"io"
	"unicode"
)

type CommentLexer struct{}

func NewCommentLexer() Lexable {
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

// Lex lexes from after '//' to end of line. It supports multi line comments.
func (c *CommentLexer) Lex(r Readable) (results []rune) {
	for c.Match(r) {
		r.ReadRunes(2) // throwaway '//' at beginning of line

		results = append(results, r.ReadTill(
			func(char rune) bool { return char != '\n' },
		)...)

		// read '\n' and add to results
		chars, err := r.ReadRunes(1)
		if err != nil {
			return results
		}

		results = append(results, chars[0])
	}

	return results
}

type NumberLexer struct{}

func NewNumberLexer() Lexable {
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

func NewStringLexer() Lexable {
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
// string and escaped \" and escaped characters inside string.
//
// TODO: raise error on unterminated strings.
func (s *StringLexer) Lex(r Readable) (results []rune) {
	r.ReadRunes(1) // throwaway " at beginning of line

	for {
		chars, err := r.ReadRunes(1)
		if err != nil || chars[0] == '"' { // end of string
			return results
		}

		results = append(results, chars[0])

		// if escape character, read next char blindly and add to results
		if chars[0] == '\\' {
			if chars, err = r.ReadRunes(1); err != nil {
				return results
			}
			results = append(results, chars[0])
		}
	}

	r.ReadRunes(1) // throwaway " at end of line

	return results
}

type IdentifierLexer struct{}

func NewIdentifierLexer() Lexable {
	return &IdentifierLexer{}
}

// Match matches if first character is a number.
func (i *IdentifierLexer) Match(p Peekable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	return unicode.IsLetter(char)
}

// Lex lexes from start till space, tab, end of line or carriage return.
func (i *IdentifierLexer) Lex(r Readable) []rune {
	return r.ReadTill(
		func(char rune) bool {
			return unicode.IsLetter(char) || unicode.IsNumber(char)
		},
	)
}

type EOFLexer struct{}

func NewEOFLexer() Lexable {
	return &EOFLexer{}
}

// Match matches if at end of input.
func (e *EOFLexer) Match(p Peekable) bool {
	_, err := p.PeekSingleRune()
	return err == io.EOF
}

// Lex returns nil to indicate there's nothing more to lex.
func (e *EOFLexer) Lex(r Readable) []rune {
	return nil
}

var WhiteSpaceChars = []string{"\t", "\n", "\r", " "}

type WhiteSpaceLexer struct{}

func NewWhiteSpaceLexer() Lexable {
	return &WhiteSpaceLexer{}
}

func (w *WhiteSpaceLexer) Match(p Peekable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	for _, whiteSpaceChar := range WhiteSpaceChars {
		if whiteSpaceChar == string(char) {
			return true
		}
	}

	return false
}

func (w *WhiteSpaceLexer) Lex(r Readable) []rune {
	chars, err := r.ReadRunes(1)
	if err != nil {
		return nil
	}

	return chars
}
