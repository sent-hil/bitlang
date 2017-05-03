package lexer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sent-hil/bitlang/runeio"
	"github.com/sent-hil/bitlang/token"
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
				So(commentRunes[0].Value, ShouldEqual, "Hello World")
			})

			Convey("It returns comments inside comments", func() {
				commentRunes := l.Lex(newRuneReader("// Hello // World"))
				So(commentRunes[0].Value, ShouldEqual, " Hello // World")
			})

			Convey("It returns multi line comments with newlines", func() {
				commentRunes := l.Lex(newRuneReader("// Hello\n// World\n"))
				So(commentRunes[0].Value, ShouldEqual, " Hello")
				So(commentRunes[1].Value, ShouldEqual, "\n")
				So(commentRunes[2].Value, ShouldEqual, " World")
				So(commentRunes[3].Value, ShouldEqual, "\n")
			})

			Convey("It does not lex anything after new line", func() {
				commentRunes := l.Lex(newRuneReader("// Hello World\n//"))
				So(len(commentRunes), ShouldEqual, 3)
				So(commentRunes[0].Value, ShouldEqual, " Hello World")
				So(commentRunes[1].Value, ShouldEqual, "\n")
				So(commentRunes[2].Value, ShouldEqual, "")
			})
		})
	})
}

func TestNumberLexer(t *testing.T) {
	Convey("NumberLexer", t, func() {
		l := NewNumberLexer()

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
				lexed := l.Lex(newRuneReader("1234"))
				So(lexed[0].Value, ShouldEqual, "1234")
				So(lexed[0].ID, ShouldEqual, token.INTEGER)
			})

			Convey("It returns till end of float", func() {
				lexed := l.Lex(newRuneReader("1234.5"))
				So(lexed[0].Value, ShouldEqual, "1234.5")
				So(lexed[0].ID, ShouldEqual, token.FLOAT)
			})

			Convey("It does not lex anything after 1st dot", func() {
				lexmes := l.Lex(newRuneReader("1234.5.6"))
				So(lexmes[0].Value, ShouldEqual, "1234.5")
				So(len(lexmes), ShouldEqual, 1)
			})

			Convey("It does not lex anything after number", func() {
				lexmes := l.Lex(newRuneReader("1234 Hello"))
				So(lexmes[0].Value, ShouldEqual, "1234")
				So(len(lexmes), ShouldEqual, 1)
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
				So(l.Lex(newRuneReader(`"Hello"`))[0].Value, ShouldEqual, "Hello")
			})

			Convey("It returns escaped double quote inside double quotes", func() {
				So(l.Lex(newRuneReader(`"He\"llo"`))[0].Value, ShouldEqual, `He\"llo`)
			})

			Convey("It returns escaped slashes and double quotes inside double quotes", func() {
				//So(string(l.Lex(newRuneReader(`"He\\\"llo"`))), ShouldEqual, `He\\\"llo`)
				So(l.Lex(newRuneReader(`"He\\\"llo"`))[0].Value, ShouldEqual, `He\\\"llo`)
			})

			Convey("It returns till end of file for unterminated strings", func() {
				So(l.Lex(newRuneReader(`"Hello`))[0].Value, ShouldEqual, "Hello")
			})

			Convey("It discards quotes at end of the string", func() {
				r := newRuneReader(`"Hello"`)
				So(l.Lex(r)[0].Value, ShouldEqual, "Hello")

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
				So(l.Lex(newRuneReader("hello"))[0].Value, ShouldEqual, "hello")
			})

			Convey("It returns chars till space", func() {
				So(l.Lex(newRuneReader("hello world"))[0].Value, ShouldEqual, "hello")
			})

			Convey("It returns chars till new tab", func() {
				So(l.Lex(newRuneReader("hello\tworld"))[0].Value, ShouldEqual, "hello")
			})

			Convey("It returns chars till end of line", func() {
				So(l.Lex(newRuneReader("hello\nworld"))[0].Value, ShouldEqual, "hello")
			})

			Convey("It returns chars till carriage return", func() {
				So(l.Lex(newRuneReader("hello\rworld"))[0].Value, ShouldEqual, "hello")
			})

			Convey("It returns chars till //", func() {
				So(l.Lex(newRuneReader("hello//"))[0].Value, ShouldEqual, "hello")
			})

			Convey("It returns chars with numbers", func() {
				So(l.Lex(newRuneReader("hello1"))[0].Value, ShouldEqual, "hello1")
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
				So(l.Lex(newRuneReader(""))[0].Value, ShouldEqual, "")
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
					So(l.Lex(newRuneReader(charWithExtra))[0].Value, ShouldEqual, char)
				}
			})
		})
	})
}
