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

				b, ok := hw.RuneReadUnreader.(*bytes.Buffer)
				So(ok, ShouldBeTrue)
				So(b.String(), ShouldEqual, "ello World")
			})

			Convey("It discards all runes when given length is same length as reader", func() {
				discarded, err := hw.Discard(11)
				So(err, ShouldEqual, nil)
				So(discarded, ShouldEqual, 11)

				b, ok := hw.RuneReadUnreader.(*bytes.Buffer)
				So(ok, ShouldBeTrue)
				So(b.String(), ShouldEqual, "")
			})

			Convey("It returns io.EOF when given length is greater than length in reader", func() {
				discarded, err := hw.Discard(12)
				So(err, ShouldEqual, io.EOF)
				So(discarded, ShouldEqual, 11)

				b, ok := hw.RuneReadUnreader.(*bytes.Buffer)
				So(ok, ShouldBeTrue)
				So(b.String(), ShouldEqual, "")

				discarded, err = em.Discard(1)
				So(err, ShouldEqual, io.EOF)
				So(discarded, ShouldEqual, 0)
			})
		})
	})
}
