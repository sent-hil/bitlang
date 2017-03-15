package lexer

type Peekable interface {
	PeekRunes(uint) ([]rune, error)
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
	return r.ReadTill(func(r rune) bool { return r != '\n' })
}
