package facet

type Facet struct {
	bitCount int
}

func New() *Facet {
	return &Facet{}
}

func (facet *Facet) Set(bit int) {
	facet.bitCount++
}

func (facet *Facet) Count() int {
	return facet.bitCount
}
