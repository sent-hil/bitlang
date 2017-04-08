package lexer

import "io"

var SymbolsMap = map[string]string{
	"(": "(",
	")": ")",
	"{": "{",
	"}": "}",
	",": ",",
	".": ".",
	"-": "-",
	"+": "+",
	";": ";",
	"/": "/",
	"!": "!",
	"=": "=",
	"<": "<",
	">": ">",
}

var SymbolsNested = map[string]string{
	"!=": "!=",
	"==": "==",
	"<=": "<",
	">=": ">=",
}

type SymbolLexer struct{}

func NewSymbolLexer() Lexable {
	return &SymbolLexer{}
}

func (s *SymbolLexer) Match(p Peekable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	_, exists := SymbolsMap[string(char)]
	return exists
}

func (s *SymbolLexer) Lex(r Readable) (result []rune) {
	chars, err := r.PeekRunes(2)
	if err != nil && err != io.EOF {
		return nil
	}
	if err == io.EOF { // reached end of line, so there only be 1 symbol, if that
		return s.LexSingle(r)
	}

	if _, ok := SymbolsNested[string(chars)]; ok {
		if result, err = r.ReadRunes(2); err != nil {
			return nil
		}
		return result
	}

	return s.LexSingle(r)
}

func (s *SymbolLexer) LexSingle(r Readable) (result []rune) {
	char, err := r.PeekRunes(1)
	if err != nil {
		return nil
	}

	if _, ok := SymbolsMap[string(char)]; ok {
		if result, err = r.ReadRunes(1); err != nil {
			return nil
		}
	}

	return result
}
