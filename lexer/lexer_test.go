package lexer

import (
	"strings"
	"testing"

	"github.com/sent-hil/bitlang/runeio"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAnyLexer(t *testing.T) {
	Convey("AnyLexer", t, func() {
		l := NewAnyLexer(runeio.NewReader(
			strings.NewReader("//comment\n1.23\nif{}\"string\" ident"),
		))

		results, err := l.LexAll()
		So(err, ShouldEqual, nil)
		So(string(results[0]), ShouldEqual, "comment\n")
		So(string(results[1]), ShouldEqual, "1.23")
		So(string(results[2]), ShouldEqual, "\n")
		So(string(results[3]), ShouldEqual, "if")
		So(string(results[4]), ShouldEqual, "{")
		So(string(results[5]), ShouldEqual, "}")
		So(string(results[6]), ShouldEqual, "string")
		So(string(results[7]), ShouldEqual, " ")
		So(string(results[8]), ShouldEqual, "ident")
	})
}
