package runeio

import (
	"bytes"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRuneIo(t *testing.T) {
	Convey("RuneIo", t, func() {
		hw := NewRuneio(bytes.NewBufferString("Hello World"))
		em := NewRuneio(bytes.NewBufferString(""))

		Convey("NewRuneio", func() {
			Convey("It returns initialized Reader", func() {
				So(hw, ShouldHaveSameTypeAs, &Reader{})
			})
		})

		Convey("Discard", func() {
			Convey("It discards given length of runes", func() {
				discarded, err := hw.Discard(1)
				So(err, ShouldEqual, nil)
				So(discarded, ShouldEqual, 1)
				So(hw.String(), ShouldEqual, "ello World")
			})

			Convey("It discards all runes when given length is same length as reader", func() {
				discarded, err := hw.Discard(11)
				So(err, ShouldEqual, nil)
				So(discarded, ShouldEqual, 11)
				So(hw.String(), ShouldEqual, "")
			})

			Convey("It returns io.EOF when given length is greater than length in reader", func() {
				discarded, err := hw.Discard(12)
				So(err, ShouldEqual, io.EOF)
				So(discarded, ShouldEqual, 11)
				So(hw.String(), ShouldEqual, "")

				discarded, err = em.Discard(1)
				So(err, ShouldEqual, io.EOF)
				So(discarded, ShouldEqual, 0)
			})
		})

		Convey("ReadRunes", func() {
			Convey("It discards given length of runes", func() {
				runes, err := hw.ReadRunes(1)
				So(err, ShouldBeNil)
				So(runes, ShouldHaveSameTypeAs, []rune{})
				So(string(runes), ShouldResemble, "H")
			})

			Convey("It returns all runes when given length is same length as reader", func() {
				runes, err := hw.ReadRunes(11)
				So(err, ShouldBeNil)
				So(string(runes), ShouldEqual, "Hello World")
			})

			Convey("It returns io.EOF when given length is greater than length in reader", func() {
				runes, err := hw.ReadRunes(12)
				So(err, ShouldEqual, io.EOF)
				So(string(runes), ShouldEqual, "Hello World")
			})

			Convey("It removes runes from reader", func() {
				_, err := hw.ReadRunes(11)
				So(err, ShouldBeNil)
				So(hw.String(), ShouldEqual, "")
			})
		})

		Convey("PeekRunes", func() {
			Convey("It returns given length of runes", func() {
				runes, err := hw.PeekRunes(1)
				So(err, ShouldBeNil)
				So(runes, ShouldHaveSameTypeAs, []rune{})
				So(string(runes), ShouldResemble, "H")
			})

			Convey("It returns io.EOF when given length is greater than length in reader", func() {
				runes, err := hw.PeekRunes(12)
				So(err, ShouldEqual, io.EOF)
				So(string(runes), ShouldEqual, "Hello World")
			})

			Convey("It does not remove runes from reader", func() {
				_, err := hw.PeekRunes(1)
				So(err, ShouldBeNil)
				So(hw.String(), ShouldEqual, "Hello World")
			})
		})
	})
}
