package lexer

import (
	"strings"
	"testing"

	"github.com/sent-hil/bitlang/runeio"
	"github.com/sent-hil/bitlang/token"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAnyLexer(t *testing.T) {
	Convey("AnyLexer", t, func() {
		l := NewAnyLexer(runeio.NewReader(
			strings.NewReader("//This is a comment\n1.23\nif{}\"string\" ident"),
		))

		results, err := l.LexAll()
		So(err, ShouldEqual, nil)

		So(results[0].Value, ShouldEqual, "This is a comment")
		So(results[0].ID, ShouldEqual, token.COMMENT)

		So(results[1].Value, ShouldEqual, "\n")
		So(results[1].ID, ShouldEqual, token.WHITESPACE)

		So(results[2].Value, ShouldEqual, "1.23")
		So(results[2].ID, ShouldEqual, token.FLOAT)

		So(results[3].Value, ShouldEqual, "\n")
		So(results[3].ID, ShouldEqual, token.WHITESPACE)

		So(results[4].Value, ShouldEqual, "if")
		So(results[4].ID, ShouldEqual, token.IF)

		So(results[5].Value, ShouldEqual, "{")
		So(results[5].ID, ShouldEqual, token.LEFT_BRACE)

		So(results[6].Value, ShouldEqual, "}")
		So(results[6].ID, ShouldEqual, token.RIGHT_BRACE)

		So(results[7].Value, ShouldEqual, "string")
		So(results[7].ID, ShouldEqual, token.STRING)

		So(results[8].Value, ShouldEqual, " ")
		So(results[8].ID, ShouldEqual, token.WHITESPACE)

		So(results[9].Value, ShouldEqual, "ident")
		So(results[9].ID, ShouldEqual, token.IDENTIFIER)
	})
}
