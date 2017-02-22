package lexer

import (
	"strings"
	"testing"

	"github.com/sent-hil/bitlang/runeio"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLexer(t *testing.T) {
	Convey("Lexer", t, func() {
		_, err := NewLexer(runeio.NewReader(strings.NewReader("Hello")))
		So(err, ShouldBeNil)
	})
}
