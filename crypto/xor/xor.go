package xor

import "io"

type xorStream struct {
	src io.Reader
	pos int
	key []byte
}

// NewReader creates new reader that for each byte of source reader applises XOR
// operation with each byte of key. Key is repeated in cycle as much as
// necessary to process all bytes of source reader.
func NewReader(src io.Reader, key ...byte) io.Reader {
	return &xorStream{src: src, key: key}
}

func (s *xorStream) Read(p []byte) (n int, err error) {
	n, err = s.src.Read(p)
	s.xor(p[:n])
	return
}

func (s *xorStream) xor(p []byte) {
	if len(s.key) == 0 {
		return
	}
	for i := 0; i < len(p); i, s.pos = i+1, s.pos+1 {
		p[i] ^= s.key[s.pos%len(s.key)]
	}
}
