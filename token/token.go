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
	COMMENT
	WHITESPACE
	STRING
	NIL
	FLOAT
	INTEGER
	EOF // THIS NEEDS TO BE LAST ONE IN LIST FOR CHECKS
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
	COMMENT:       "COMMENT",
	WHITESPACE:    "WHITESPACE",
	STRING:        "STRING",
	NIL:           "NIL",
	FLOAT:         "FLOAT",
	INTEGER:       "INTEGER",
	EOF:           "EOF",
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

func init() {
	// check to make sure all tokens are in TokenIDString list; for this to
	// properly work EOF needs to be the last token defined in iota above.
	for i := 0; i <= int(EOF); i++ {
		if _, ok := TokenIDString[TokenID(i)]; !ok {
			panic(fmt.Sprintf("Missing token: %v in TokenIDString list", i))
		}
	}
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
