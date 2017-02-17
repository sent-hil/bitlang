package runeio

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRuneIo(t *testing.T) {
	r := NewRuneio(bytes.NewBufferString("Hello World"))

	Convey("NewRuneio", t, func() {
		Convey("It returns initialized Reader", func() {
			So(r, ShouldHaveSameTypeAs, &Reader{})
		})
	})
}
