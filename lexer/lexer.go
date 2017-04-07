package lexer

import "github.com/sent-hil/bitlang/runeio"

type Peekable interface {
	PeekRunes(uint) ([]rune, error)
	PeekSingleRune() (rune, error)
}

type Readable interface {
	Peekable
	ReadRunes(uint) ([]rune, error)
	ReadTill(func(rune) bool) []rune
}

type Lexable interface {
	Match(p Peekable) bool
	Lex(r Readable) []rune
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

func (a *AnyLexer) LexAll() (results [][]rune, err error) {
	for !a.reader.IsAtEnd() {
		for _, initializer := range a.lexers {
			lexer := initializer()
			if lexer.Match(a.reader) {
				results = append(results, lexer.Lex(a.reader))
				break
			}
		}
	}

	return results, err
}
