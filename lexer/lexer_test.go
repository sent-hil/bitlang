package lexer

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLexer(t *testing.T) {
	Convey("NewLexer", t, func() {
		Convey("It parses and stores characters to be iterated", func() {
			l, err := NewLexer(strings.NewReader("Hello"))
			So(err, ShouldBeNil)
			So(len(l.runes), ShouldEqual, 5)
			So(l.runes, ShouldResemble, []rune{'H', 'e', 'l', 'l', 'o'})
		})
	})
}
