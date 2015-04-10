package stats

import "io"

// CodeStats represents frequencies of bytes.
type CodeStats struct {
	counts [256]int
	buf    []byte
}

// New creates new zero code stats with buffer provided.
//
// Buffer is used for reading from reader later. We allow to provide buffer to
// avoid repeatable allocations since stats are recalculated frequently.
func New(buf []byte) CodeStats {
	return CodeStats{buf: buf}
}

// Reset resets all counts to 0.
func (s *CodeStats) Reset() {
	s.counts = [256]int{}
}

// ReadFrom implements io.ReadFrom interface. It read all bytes from the reader
// until EOF and increments these bytes counts.
func (s *CodeStats) ReadFrom(r io.Reader) (n int64, err error) {
	for err == nil {
		var m int
		m, err = r.Read(s.buf)
		s.Write(s.buf[:m])
		n += int64(m)
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func (s *CodeStats) Write(p []byte) (n int, err error) {
	for _, b := range p {
		s.counts[b]++
	}
	return len(p), nil
}

// English returns known stats for an english text.
func English() CodeStats {
	e := CodeStats{
		counts: [256]int{
			'a': 8167,
			'b': 1492,
			'c': 2782,
			'd': 4253,
			'e': 12702,
			'f': 2228,
			'g': 2015,
			'h': 6094,
			'i': 6966,
			'j': 153,
			'k': 772,
			'l': 4025,
			'm': 2406,
			'n': 6749,
			'o': 7507,
			'p': 1929,
			'q': 95,
			'r': 5987,
			's': 6327,
			't': 9056,
			'u': 2758,
			'v': 978,
			'w': 2360,
			'x': 150,
			'y': 1974,
			'z': 74,
		},
	}
	return e
}

// D returns square of dispersion of each byte frequencies in stats.
func (s CodeStats) D() float64 {
	return Mul(s, s)
}

func (s CodeStats) total() int {
	t := 0
	for _, c := range s.counts {
		t += c
	}
	return t
}

// Mul does vector multiplication of each byte frequencies of two stats.
func Mul(s1, s2 CodeStats) float64 {
	t1, t2 := s1.total(), s2.total()
	if t1 == 0 || t2 == 0 {
		return 0
	}
	var m float64
	for i := 0; i < 256; i++ {
		m += (float64(s1.counts[i]) / float64(t1)) * (float64(s2.counts[i]) / float64(t2))
	}
	return m
}
