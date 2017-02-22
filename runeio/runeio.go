package runeio

import (
	"io"
	"io/ioutil"
)

// RuneReader is the underlying interface Reader will use.
type RuneReader interface {
	ReadRune() (r rune, size int, err error)
	io.Reader
}

type Reader struct {
	RuneReader
	Runes []rune
}

func NewRuneio(r RuneReader) *Reader {
	return &Reader{r, []rune{}}
}

func (r *Reader) Discard(n uint) (uint, error) {
	runes, err := r.ReadRunes(n)
	return uint(len(runes)), err
}

func (r *Reader) ReadRunes(n uint) (runes []rune, err error) {
	if err = r.readFromReader(n); err != nil {
		n = uint(len(r.Runes))
	}

	runes = r.Runes[0:n]
	r.Runes = r.Runes[n:]

	return runes, err
}

func (r *Reader) PeekRunes(n uint) ([]rune, error) {
	if err := r.readFromReader(n); err != nil {
		return r.Runes, err
	}

	return r.Runes[0:n], nil
}

func (r *Reader) PeekRune() (rune, error) {
	runes, err := r.PeekRunes(1)
	if err != nil {
		return '\uFFFD', err
	}

	return runes[0], nil
}

func (r *Reader) String() (string, error) {
	bites, err := ioutil.ReadAll(r.RuneReader)
	if err != nil {
		return "", err
	}
	return string(r.Runes) + string(bites), nil
}

func (r *Reader) Reset(bufReader RuneReader) {
	r.RuneReader = bufReader
}

func (r *Reader) readFromReader(n uint) error {
	l := int(n) - len(r.Runes)

	// check if we've already read enough runes
	if l <= 0 {
		return nil
	}

	// if not, read the remaining amount of runes
	for i := 0; i < l; i++ {
		ru, _, err := r.ReadRune()
		if err != nil {
			return err
		}
		r.Runes = append(r.Runes, ru)
	}

	return nil
}
