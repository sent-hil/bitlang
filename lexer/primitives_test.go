package lexer

import (
	"bytes"
	"testing"

	"github.com/sent-hil/bitlang/runeio"
	. "github.com/smartystreets/goconvey/convey"
)

func newRuneReader(s string) *runeio.Reader {
	return runeio.NewReader(bytes.NewBufferString(s))
}

func TestPrimitives(t *testing.T) {
	Convey("CommentLexer", t, func() {
		Convey("Match", func() {
			Convey("It does not match if only '/' is present", func() {
				So(NewCommentLexer().Match(newRuneReader("/")), ShouldEqual, false)
			})

			Convey("It does not match on empty string", func() {
				So(NewCommentLexer().Match(newRuneReader("")), ShouldEqual, false)
			})

			Convey("It matches if '//' are in beginning of line", func() {
				So(NewCommentLexer().Match(newRuneReader("//")), ShouldEqual, true)
			})
		})

		Convey("Lex", func() {
			cLexer := NewCommentLexer()

			Convey("It returns string till end of line", func() {
				commentRunes := cLexer.Lex(newRuneReader("// Hello World"))
				So(string(commentRunes), ShouldEqual, "// Hello World")
			})

			Convey("It should not lex anything after new line", func() {
				commentRunes := cLexer.Lex(newRuneReader("// Hello World\n//"))
				So(string(commentRunes), ShouldEqual, "// Hello World")
			})
		})
	})

	Convey("IntegerLexer", t, func() {
		l := NewIntegerLexer()

		Convey("Match", func() {
			Convey("It matches if char is a number", func() {
				So(l.Match(newRuneReader("1")), ShouldEqual, true)
			})

			Convey("It does not match on empty string", func() {
				So(l.Match(newRuneReader("")), ShouldEqual, false)
			})

			Convey("It does not match on strings", func() {
				So(l.Match(newRuneReader("Hello")), ShouldEqual, false)
			})
		})

		Convey("Lex", func() {
			Convey("It returns till end of integer", func() {
				So(string(l.Lex(newRuneReader("1234"))), ShouldEqual, "1234")
			})

			Convey("It should not lex anything after integer", func() {
				So(string(l.Lex(newRuneReader("1234 Hello"))), ShouldEqual, "1234")
			})
		})
	})
}
