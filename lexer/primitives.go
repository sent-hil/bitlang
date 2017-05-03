package lexer

import (
	"io"
	"unicode"

	"github.com/sent-hil/bitlang/token"
)

type CommentLexer struct{}

func NewCommentLexer() Lexable {
	return &CommentLexer{}
}

// Match matches if 1st and 2nd characters are /, ie //.
func (c *CommentLexer) Match(p Readable) bool {
	chars, err := p.PeekRunes(2)
	if err != nil {
		return false
	}

	return string(chars) == "//"
}

// Lex lexes from after // to end of line. It can parse multi line comments,
// but each line needs to be prefixed with //.
func (c *CommentLexer) Lex(r Readable) (tokens []*token.Token) {
	for c.Match(r) {
		r.ReadRunes(2) // throwaway '//' at beginning of line

		singleLine := r.ReadTill(
			func(char rune) bool { return char != '\n' },
		)

		tokens = append(tokens, token.NewToken(token.COMMENT, string(singleLine)))

		// read '\n' at end of line and add to tokens
		singleLine, err := r.ReadRunes(1)
		if err != nil {
			return tokens
		}

		tokens = append(tokens,
			token.NewToken(token.WHITESPACE, string(singleLine)))
	}

	return tokens
}

type NumberLexer struct{}

func NewNumberLexer() Lexable {
	return &NumberLexer{}
}

// Match matches if first character is a digit.
func (i *NumberLexer) Match(p Readable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	return unicode.IsNumber(char)
}

// Lex lexes integers and floats.
func (i *NumberLexer) Lex(r Readable) (tokens []*token.Token) {
	hasDot := false
	tokenId := token.INTEGER

	accum := r.ReadTill(
		func(char rune) bool {
			if unicode.IsNumber(char) {
				return true
			}

			if char == '.' {
				if hasDot {
					return false // already has a dot, so this is a method call
				} else {
					tokenId = token.FLOAT
					hasDot = true
					return true
				}
			}

			return false
		},
	)

	tokens = append(tokens, token.NewToken(tokenId, string(accum)))

	return tokens
}

type StringLexer struct{}

func NewStringLexer() Lexable {
	return &StringLexer{}
}

// Match matches if first character is double quotes.
func (s *StringLexer) Match(p Readable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	return char == '"'
}

// Lex lexes all characters inside double quotes. It works with multiple line
// string and also escaped \" and escaped characters inside string.
//
// TODO: raise error on unterminated strings.
func (s *StringLexer) Lex(r Readable) (tokens []*token.Token) {
	r.ReadRunes(1) // throwaway " at beginning of line

	var accum string
	for {
		chars, err := r.ReadRunes(1)
		if err != nil || chars[0] == '"' { // end of string
			return []*token.Token{token.NewToken(token.STRING, accum)}
		}

		accum += string(chars[0])

		// if escape character, read next char blindly and add to tokens
		if string(chars[0]) == `\` {
			// if err, ie EOF, then return everything seen so far
			if chars, err = r.ReadRunes(1); err != nil {
				return []*token.Token{token.NewToken(token.STRING, accum)}
			}
			accum += string(chars[0])
		}
	}

	r.ReadRunes(1) // throwaway " at end of line

	return []*token.Token{token.NewToken(token.STRING, accum)}
}

type IdentifierLexer struct{}

func NewIdentifierLexer() Lexable {
	return &IdentifierLexer{}
}

// Match matches if first character is a number.
func (i *IdentifierLexer) Match(p Readable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	return unicode.IsLetter(char)
}

// Lex lexes from start till space, tab, end of line or carriage return.
func (i *IdentifierLexer) Lex(r Readable) []*token.Token {
	accum := r.ReadTill(
		func(char rune) bool {
			return unicode.IsLetter(char) || unicode.IsNumber(char)
		},
	)

	tId, ok := token.KeywordsList[string(accum)]
	if ok {
		return []*token.Token{token.NewToken(tId, string(accum))}
	}

	// if not a reserved keyword, then it's a identifier
	return []*token.Token{token.NewToken(token.IDENTIFIER, string(accum))}
}

type EOFLexer struct{}

func NewEOFLexer() Lexable {
	return &EOFLexer{}
}

// Match matches if at end of string.
func (e *EOFLexer) Match(p Readable) bool {
	_, err := p.PeekSingleRune()
	return err == io.EOF
}

// Lex returns nil to indicate there's nothing more to lex.
func (e *EOFLexer) Lex(r Readable) []*token.Token {
	return []*token.Token{token.NewToken(token.EOF, "")}
}

var WhiteSpaceChars = []string{"\t", "\n", "\r", " "}

type WhiteSpaceLexer struct{}

func NewWhiteSpaceLexer() Lexable {
	return &WhiteSpaceLexer{}
}

func (w *WhiteSpaceLexer) Match(p Readable) bool {
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

func (w *WhiteSpaceLexer) Lex(r Readable) []*token.Token {
	accum, err := r.ReadRunes(1)
	if err != nil {
		return nil
	}

	return []*token.Token{token.NewToken(token.WHITESPACE, string(accum))}
}
