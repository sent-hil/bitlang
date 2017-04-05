package lexer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sent-hil/bitlang/runeio"
	. "github.com/smartystreets/goconvey/convey"
)

func newRuneReader(s string) *runeio.Reader {
	return runeio.NewReader(bytes.NewBufferString(s))
}

func TestCommentLexer(t *testing.T) {
	Convey("CommentLexer", t, func() {
		l := NewCommentLexer()

		Convey("Match", func() {
			Convey("It does not match if only '/' is present", func() {
				So(l.Match(newRuneReader("/")), ShouldEqual, false)
			})

			Convey("It does not match on empty string", func() {
				So(l.Match(newRuneReader("")), ShouldEqual, false)
			})

			Convey("It matches if '//' are in beginning of line", func() {
				So(l.Match(newRuneReader("//")), ShouldEqual, true)
			})

			Convey("It matches if '//' is in beginning of line followed any char", func() {
				So(l.Match(newRuneReader("//H")), ShouldEqual, true)
			})
		})

		Convey("Lex", func() {
			Convey("It returns comments from '//' till end of line", func() {
				commentRunes := l.Lex(newRuneReader("//Hello World"))
				So(string(commentRunes), ShouldEqual, "Hello World")
			})

			Convey("It returns comments inside comments", func() {
				commentRunes := l.Lex(newRuneReader("// Hello // World"))
				So(string(commentRunes), ShouldEqual, " Hello // World")
			})

			Convey("It does not lex anything after new line", func() {
				commentRunes := l.Lex(newRuneReader("// Hello World\n//"))
				So(string(commentRunes), ShouldEqual, " Hello World")
			})
		})
	})
}

func TestNumberLexer(t *testing.T) {
	Convey("NumberLexer", t, func() {
		l := NewIntegerLexer()

		Convey("Match", func() {
			Convey("It matches if char is a number", func() {
				So(l.Match(newRuneReader("1")), ShouldEqual, true)
			})

			Convey("It does not match on empty string", func() {
				So(l.Match(newRuneReader("")), ShouldEqual, false)
			})

			Convey("It does not match any chars", func() {
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

			Convey("IT does not lex anything after 1st dot", func() {
				So(string(l.Lex(newRuneReader("1234.5.6"))), ShouldEqual, "1234.5")
			})

			Convey("It does not lex anything after number", func() {
				So(string(l.Lex(newRuneReader("1234 Hello"))), ShouldEqual, "1234")
			})
		})
	})
}

func TestStringLexer(t *testing.T) {
	Convey("StringLexer", t, func() {
		l := NewStringLexer()

		Convey("Match", func() {
			Convey("It matches if 1st char is double quotes", func() {
				So(l.Match(newRuneReader(`"`)), ShouldEqual, true)
			})

			Convey("It matches chars wrapped in double quotes", func() {
				So(l.Match(newRuneReader(`"Hello"`)), ShouldEqual, true)
			})

			Convey("It does not match single quote", func() {
				So(l.Match(newRuneReader(`'`)), ShouldEqual, false)
			})

			Convey("It does not match on empty string", func() {
				So(l.Match(newRuneReader("")), ShouldEqual, false)
			})
		})

		Convey("Lex", func() {
			Convey("It returns chars inside double quotes", func() {
				So(string(l.Lex(newRuneReader(`"Hello"`))), ShouldEqual, "Hello")
			})

			Convey("It discards quotes at end of the string", func() {
				r := newRuneReader(`"Hello"`)
				So(string(l.Lex(r)), ShouldEqual, "Hello")

				remaining, err := r.String()
				So(err, ShouldBeNil)
				So(remaining, ShouldEqual, "")
			})
		})
	})
}

func TestIdentifierLexer(t *testing.T) {
	Convey("IdentifierLexer", t, func() {
		l := NewIdentifierLexer()

		Convey("Match", func() {
			Convey("It matches if 1st char is double quotes", func() {
				So(l.Match(newRuneReader("hello")), ShouldEqual, true)
			})
		})

		Convey("Lex", func() {
			Convey("It returns chars till end of string", func() {
				So(string(l.Lex(newRuneReader("hello"))), ShouldEqual, "hello")
			})

			Convey("It returns chars till space", func() {
				So(string(l.Lex(newRuneReader("hello world"))), ShouldEqual, "hello")
			})

			Convey("It returns chars till new tab", func() {
				So(string(l.Lex(newRuneReader("hello\tworld"))), ShouldEqual, "hello")
			})

			Convey("It returns chars till end of line", func() {
				So(string(l.Lex(newRuneReader("hello\nworld"))), ShouldEqual, "hello")
			})

			Convey("It returns chars till carriage return", func() {
				So(string(l.Lex(newRuneReader("hello\rworld"))), ShouldEqual, "hello")
			})

			Convey("It returns chars till //", func() {
				So(string(l.Lex(newRuneReader("hello//"))), ShouldEqual, "hello")
			})

			Convey("It returns chars with numbers", func() {
				So(string(l.Lex(newRuneReader("hello1"))), ShouldEqual, "hello1")
			})
		})
	})
}

func TestEOFLexer(t *testing.T) {
	Convey("EOFLexer", t, func() {
		l := NewEOFLexer()

		Convey("Match", func() {
			Convey("It matches if at end of file", func() {
				So(l.Match(newRuneReader("")), ShouldEqual, true)
			})
		})

		Convey("Lex", func() {
			Convey("It returns chars till end of string", func() {
				So(len(l.Lex(newRuneReader(""))), ShouldEqual, 0)
			})
		})
	})
}

func TestWhiteSpaceLexer(t *testing.T) {
	Convey("WhiteSpaceLexer", t, func() {
		l := NewWhiteSpaceLexer()

		Convey("Match", func() {
			Convey("It matches if at end of file", func() {
				for _, char := range WhiteSpaceChars {
					So(l.Match(newRuneReader(char)), ShouldEqual, true)
				}
			})
		})

		Convey("Lex", func() {
			Convey("It returns chars till end of string", func() {
				for _, char := range WhiteSpaceChars {
					charWithExtra := fmt.Sprintf("%shello", char)
					So(string(l.Lex(newRuneReader(charWithExtra))), ShouldEqual, char)
				}
			})
		})
	})
}
