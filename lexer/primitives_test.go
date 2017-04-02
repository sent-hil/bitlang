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
			l := NewCommentLexer()

			Convey("It returns string till end of line", func() {
				commentRunes := l.Lex(newRuneReader("//Hello World"))
				So(string(commentRunes), ShouldEqual, "Hello World")
			})

			Convey("It returns comments inside comments", func() {
				commentRunes := l.Lex(newRuneReader("// Hello // World"))
				So(string(commentRunes), ShouldEqual, " Hello // World")
			})

			Convey("It should not lex anything after new line", func() {
				commentRunes := l.Lex(newRuneReader("// Hello World\n//"))
				So(string(commentRunes), ShouldEqual, " Hello World")
			})
		})
	})

	Convey("NumberLexer", t, func() {
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

			Convey("It returns till end of float", func() {
				So(string(l.Lex(newRuneReader("1234.5"))), ShouldEqual, "1234.5")
			})

			Convey("It returns till end of float even with multiple dots", func() {
				So(string(l.Lex(newRuneReader("1234.5.6"))), ShouldEqual, "1234.5")
			})

			Convey("It should not lex anything after number", func() {
				So(string(l.Lex(newRuneReader("1234 Hello"))), ShouldEqual, "1234")
			})
		})
	})
}
