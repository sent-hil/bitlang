package token

type TokenID int

const (
	TokenERR TokenID = iota
	TokenEOF
	TokenInteger
	TokenComment
)

type Token struct {
	ID    TokenID
	Value string
}

func NewToken(id TokenID, value string) *Token {
	return &Token{ID: id, Value: value}
}
