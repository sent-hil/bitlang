package lexer

import (
	"io"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLexer(t *testing.T) {
	Convey("Lexer", t, func() {
		l, err := NewLexer(strings.NewReader("Hello"))
		So(err, ShouldBeNil)

		Convey("NewLexer", func() {
			Convey("It parses and stores chars to be iterated", func() {
				So(len(l.runes), ShouldEqual, 5)
				So(l.runes, ShouldResemble, []rune{'H', 'e', 'l', 'l', 'o'})
			})
		})

		Convey("Peak", func() {
			Convey("It peeks into stored chars when given size is within stored chars size", func() {
				chars, err := l.PeekRunes(2)
				So(err, ShouldBeNil)
				So(chars, ShouldResemble, []rune{'H', 'e'})
			})

			Convey("It peeks entire stored chars slice", func() {
				chars, err := l.PeekRunes(5)
				So(err, ShouldBeNil)
				So(chars, ShouldResemble, []rune{'H', 'e', 'l', 'l', 'o'})
			})

			Convey("It returns error when give size is greater than stored chars size", func() {
				_, err := l.PeekRunes(6)
				So(err, ShouldEqual, ErrBoundsExceeded)
			})

			Convey("It returns error when give size is negative", func() {
				_, err := l.PeekRunes(-1)
				So(err, ShouldEqual, ErrNegativeCount)
			})
		})

		Convey("ReadRune", func() {
			Convey("It returns chars as they're read from chars slice", func() {
				// required since l.Read modifies l.runes
				dupRunes := make([]rune, len(l.runes))
				copy(dupRunes, l.runes)

				for i := 0; i < len(dupRunes); i++ {
					char, err := l.ReadRune()
					So(err, ShouldBeNil)

					// cast to string for readable errors
					So(string(char), ShouldEqual, string(dupRunes[i]))
				}

				Convey("It returns EOF if no chars are left in chars slice", func() {
					char, err := l.ReadRune()
					So(err, ShouldEqual, io.EOF)
					So(char, ShouldEqual, '\uFFFD') // error rune
				})
			})
		})

		Convey("ReadRunes", func() {
			Convey("It returns given size of chars when size is within stored chars size", func() {
				chars, err := l.ReadRunes(1)
				So(err, ShouldBeNil)
				So(chars, ShouldResemble, []rune{'H'})
			})

			Convey("It returns entire stored chars slice", func() {
				chars, err := l.ReadRunes(5)
				So(err, ShouldBeNil)
				So(chars, ShouldResemble, []rune{'H', 'e', 'l', 'l', 'o'})
			})

			Convey("It removes returned chars from stored chars slice", func() {
				_, err := l.ReadRunes(1)
				So(err, ShouldBeNil)
				So(l.runes, ShouldResemble, []rune{'e', 'l', 'l', 'o'})

				chars, err := l.ReadRunes(1)
				So(err, ShouldBeNil)
				So(chars, ShouldResemble, []rune{'e'})
			})

			Convey("It returns error when give size is greater than stored chars size", func() {
				_, err := l.ReadRunes(6)
				So(err, ShouldEqual, ErrBoundsExceeded)
			})

			Convey("It returns error when give size is negative", func() {
				_, err := l.ReadRunes(-1)
				So(err, ShouldEqual, ErrNegativeCount)
			})
		})
	})
}
