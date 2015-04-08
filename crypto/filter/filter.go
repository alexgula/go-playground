package filter

import "io"

type periodFilter struct {
	src           io.Reader
	pos           int
	start, period int
}

// NewReader creaes new reader from existing reader, that will skip start amount
// of bytes and then will return each period's byte.
//
// Example: NewReader(_, 2, 2) applied to reader returning [0 1 2 3 4 5] will
// return reader returning [2 4].
func NewReader(src io.Reader, start, period int) io.Reader {
	return &periodFilter{src: src, start: start, period: period}
}

func (f *periodFilter) Read(p []byte) (n int, err error) {
	n, err = f.src.Read(p)
	n = f.filter(p[:n])
	return
}

func (f *periodFilter) filter(p []byte) (n int) {
	for i := f.start; i < len(p); i, n = i+f.period, n+1 {
		p[n] = p[i]
	}

	// Skip first bytes on first run only
	f.start -= len(p)
	if f.start < 0 {
		f.start = 0
	}

	return
}
