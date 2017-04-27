package lexer

import (
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

func (a *AnyLexer) LexAll() (results []*token.Token, err error) {
	for !a.reader.IsAtEnd() {
		for _, initializer := range a.lexers {
			lexer := initializer()
			if lexer.Match(a.reader) {
				results = append(results, lexer.Lex(a.reader)...)
				break
			}
		}
	}

	return results, err
}
