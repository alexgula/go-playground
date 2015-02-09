package facet

type Facet struct {
	bits byte
}

func New() *Facet {
	return &Facet{}
}

func (f *Facet) Set(bit uint64) {
	f.bits = f.bits | (1 << bit)
}

func (f *Facet) Count() uint64 {
	n := uint64(0)
	bits := f.bits
	for i := 0; i < 8; i++ {
		if bits&1 == 1 {
			n++
		}
		bits >>= 1
	}
	return n
}
