package lexer

import (
	"fmt"

	"github.com/sent-hil/bitlang/runeio"
	"github.com/sent-hil/bitlang/token"
)

type Readable interface {
	PeekRunes(uint) ([]rune, error)
	PeekSingleRune() (rune, error)
	ReadRunes(uint) ([]rune, error)
	ReadTill(func(rune) bool) []rune
}

type Lexable interface {
	Match(p Readable) bool
	Lex(r Readable) []*token.Token
}

type LexableConstructor func() Lexable

type AnyLexer struct {
	reader *runeio.Reader
	lexers []LexableConstructor
}

func NewAnyLexer(reader *runeio.Reader) *AnyLexer {
	return &AnyLexer{
		reader: reader,
		lexers: []LexableConstructor{
			NewCommentLexer,
			NewNumberLexer,
			NewWhiteSpaceLexer,
			NewIdentifierLexer,
			NewSymbolLexer,
			NewStringLexer,
			NewEOFLexer,
		},
	}
}

func (a *AnyLexer) LexAll() (tokens []*token.Token, err error) {
	for !a.reader.IsAtEnd() {
		unMatched := true
		for _, lexerInitializer := range a.lexers {
			if lexer := lexerInitializer(); lexer.Match(a.reader) {
				unMatched = false
				tokens = append(tokens, lexer.Lex(a.reader)...)
				break
			}
		}

		if unMatched {
			char, err := a.reader.PeekSingleRune()
			if err == nil {
				return tokens, fmt.Errorf("Unmatched char: %s", char)
			}
		}
	}

	return tokens, err
}
