package runeio

// RuneReadUnreader is the underlying interface Reader will use.
type RuneReadUnreader interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
	String() string
}

type Reader struct {
	RuneReadUnreader
}

func NewRuneio(r RuneReadUnreader) *Reader {
	return &Reader{r}
}

func (r *Reader) Discard(n uint) (discarded uint, err error) {
	for i := 0; i < int(n); i++ {
		_, size, err := r.ReadRune() // size will always be 0 if there's an error
		if err != nil {
			return discarded, err
		}
		discarded += uint(size)
	}

	return discarded, nil
}

func (r *Reader) ReadRunes(n uint) (runes []rune, err error) {
	for i := 0; i < int(n); i++ {
		r, _, err := r.ReadRune()
		if err != nil {
			return runes, err
		}
		runes = append(runes, r)
	}

	return runes, nil
}
