package lexer

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSymbolLexer(t *testing.T) {
	Convey("SymbolLexer", t, func() {
		s := NewSymbolLexer()

		Convey("Match", func() {
			Convey("It matches symbol in list", func() {
				for key, _ := range SymbolsMap {
					So(s.Match(newRuneReader(key)), ShouldEqual, true)
				}
			})
		})

		Convey("Lex", func() {
			Convey("It lexes single symbol character", func() {
				for key, _ := range SymbolsMap {
					var char string

					char = fmt.Sprintf("%s", key)
					So(s.Lex(newRuneReader(char))[0].Value, ShouldEqual, key)

					char = fmt.Sprintf("%s hello", key)
					So(s.Lex(newRuneReader(char))[0].Value, ShouldEqual, key)
				}
			})

			Convey("It lexes double symbol characters", func() {
				for key, _ := range SymbolsNested {
					var char string

					char = fmt.Sprintf("%s", key)
					So(s.Lex(newRuneReader(char))[0].Value, ShouldEqual, key)

					char = fmt.Sprintf("%s hello", key)
					So(s.Lex(newRuneReader(char))[0].Value, ShouldEqual, key)
				}
			})

			Convey("It lexes double symbols chars first", func() {
				So(s.Lex(newRuneReader("!="))[0].Value, ShouldEqual, "!=")
				So(s.Lex(newRuneReader("!=="))[0].Value, ShouldEqual, "!=")
				So(s.Lex(newRuneReader("="))[0].Value, ShouldEqual, "=")
			})
		})
	})
}
