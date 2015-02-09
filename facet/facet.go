package facet

type Facet struct {
	bits byte
}

func New() *Facet {
	return &Facet{}
}

func (facet *Facet) Set(bit uint64) {
	facet.bits = facet.bits | (1 << bit)
}

func (facet *Facet) Count() (count uint64) {
	bits := facet.bits
	for i := 0; i < 8; i++ {
		if bits&1 == 1 {
			count++
		}
		bits >>= 1
	}
	return
}
