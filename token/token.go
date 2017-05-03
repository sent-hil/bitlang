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
	AND
	IF
	ELSE
	TRUE
	FALSE
	FOR
	OR
	RETURN
	VAR
	EOF
	COMMENT
	WHITESPACE
	STRING
	NUMBER
	NIL
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
	AND:           "AND",
	IF:            "IF",
	ELSE:          "ELSE",
	TRUE:          "TRUE",
	FALSE:         "FALSE",
	FOR:           "FOR",
	OR:            "OR",
	RETURN:        "RETURN",
	VAR:           "VAR",
	EOF:           "EOF",
	COMMENT:       "COMMENT",
	WHITESPACE:    "WHITESPACE",
	STRING:        "STRING",
	NUMBER:        "NUMBER",
	NIL:           "NIL",
}

var IdentifiersList = map[string]TokenID{
	"and":    AND,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"for":    FOR,
	"nil":    NIL,
	"or":     OR,
	"return": RETURN,
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
