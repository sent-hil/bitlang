package token

import "fmt"

type TokenID int

const (
	LEFT_PAREN TokenID = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER
	STRING
	NUMBER
	AND
	ELSE
	FALSE
	FOR
	IF
	NIL
	OR
	RETURN
	TRUE
	VAR
	EOF
	COMMENT
	WHITESPACE
)

var TokenIDString = map[TokenID]string{
	LEFT_PAREN:    "LEFT_PAREN",
	RIGHT_PAREN:   "RIGHT_PAREN",
	LEFT_BRACE:    "LEFT_BRACE",
	RIGHT_BRACE:   "RIGHT_BRACE",
	COMMA:         "COMMA",
	DOT:           "DOT",
	MINUS:         "MINUS",
	PLUS:          "PLUS",
	SEMICOLON:     "SEMICOLON",
	SLASH:         "SLASH",
	STAR:          "STAR",
	BANG:          "BANG",
	BANG_EQUAL:    "BANG_EQUAL",
	EQUAL:         "EQUAL",
	EQUAL_EQUAL:   "EQUAL_EQUAL",
	GREATER:       "GREATER",
	GREATER_EQUAL: "GREATER_EQUAL",
	LESS:          "LESS",
	LESS_EQUAL:    "LESS_EQUAL",
	IDENTIFIER:    "IDENTIFIER",
	STRING:        "STRING",
	NUMBER:        "NUMBER",
	AND:           "AND",
	ELSE:          "ELSE",
	FALSE:         "FALSE",
	FOR:           "FOR",
	IF:            "IF",
	NIL:           "NIL",
	OR:            "OR",
	RETURN:        "RETURN",
	TRUE:          "TRUE",
	VAR:           "VAR",
	EOF:           "EOF",
	COMMENT:       "COMMENT",
	WHITESPACE:    "WHITESPACE",
}

var IdentifiersList = map[string]TokenID{
	"and":    AND,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"return": RETURN,
	"true":   TRUE,
	"var":    VAR,
}

type Token struct {
	ID    TokenID
	Value string
}

func NewToken(id TokenID, value string) *Token {
	return &Token{ID: id, Value: value}
}

func (t *Token) String() string {
	return fmt.Sprintf("[%s] %s", TokenIDString[t.ID], t.Value)
}
