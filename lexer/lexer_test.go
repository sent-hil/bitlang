package lexer

import (
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
				chars, err := l.Peek(2)
				So(err, ShouldBeNil)
				So(chars, ShouldResemble, []rune{'H', 'e'})
			})

			Convey("It peeks entire stored chars array", func() {
				chars, err := l.Peek(5)
				So(err, ShouldBeNil)
				So(chars, ShouldResemble, []rune{'H', 'e', 'l', 'l', 'o'})
			})

			Convey("It returns error when give size is greater than stored chars size", func() {
				_, err := l.Peek(6)
				So(err, ShouldEqual, ErrBoundsExceeded)
			})

			Convey("It returns error when give size is negative", func() {
				_, err := l.Peek(-1)
				So(err, ShouldEqual, ErrNegativeCount)
			})
		})
	})
}
