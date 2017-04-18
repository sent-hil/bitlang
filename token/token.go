package token

import "fmt"

type TokenID int

const (
	TokenERR TokenID = iota
	TokenEOF
	TokenInteger
	TokenComment
	TokenNewLine
)

var IDtoString = map[TokenID]string{
	TokenERR:     "Error",
	TokenEOF:     "EOF",
	TokenInteger: "Integer",
	TokenComment: "Comment",
	TokenNewLine: "NewLine",
}

type Token struct {
	ID    TokenID
	Value string
}

func NewToken(id TokenID, value string) *Token {
	return &Token{ID: id, Value: value}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s", IDtoString[t.ID], t.Value)
}
