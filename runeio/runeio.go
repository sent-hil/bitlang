package runeio

// RuneReadUnreader is the underlying interface Reader will use.
type RuneReadUnreader interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
}

type Reader struct {
	r RuneReadUnreader
}

func NewRuneio(r RuneReadUnreader) *Reader {
	return &Reader{r: r}
}
