package lexer

import (
	"io"

	"github.com/sent-hil/bitlang/token"
)

var SymbolsMap = map[string]token.TokenID{
	"(": token.LEFT_PAREN,
	")": token.RIGHT_PAREN,
	"{": token.LEFT_BRACE,
	"}": token.RIGHT_BRACE,
	",": token.COMMA,
	".": token.DOT,
	"-": token.MINUS,
	"+": token.PLUS,
	";": token.SEMICOLON,
	"/": token.SLASH,
	"!": token.BANG,
	"=": token.EQUAL,
	"<": token.LESS,
	">": token.GREATER,
}

var SymbolsNested = map[string]token.TokenID{
	"!=": token.BANG_EQUAL,
	"==": token.EQUAL_EQUAL,
	"<=": token.LESS_EQUAL,
	">=": token.GREATER_EQUAL,
}

type SymbolLexer struct{}

func NewSymbolLexer() Lexable {
	return &SymbolLexer{}
}

func (s *SymbolLexer) Match(p Readable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	_, exists := SymbolsMap[string(char)]
	return exists
}

func (s *SymbolLexer) Lex(r Readable) (accum []*token.Token) {
	chars, err := r.PeekRunes(2)
	if err != nil && err != io.EOF {
		return nil
	}
	if err == io.EOF { // reached end of line, so there only be 1 symbol, if that
		return s.LexSingle(r)
	}

	if tId, ok := SymbolsNested[string(chars)]; ok {
		if chars, err = r.ReadRunes(2); err != nil {
			return nil
		}
		return []*token.Token{token.NewToken(tId, string(chars))}
	}

	return s.LexSingle(r)
}

func (s *SymbolLexer) LexSingle(r Readable) (accum []*token.Token) {
	char, err := r.PeekRunes(1)
	if err != nil {
		return nil
	}

	tId, ok := SymbolsMap[string(char)]
	if ok {
		if _, err := r.ReadRunes(1); err != nil {
			return nil
		}
	}

	return []*token.Token{token.NewToken(tId, string(char))}
}
