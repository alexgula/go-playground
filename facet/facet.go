package facet

type Facet struct {
	bits []byte
}

func New() *Facet {
	return &Facet{}
}

func (facet *Facet) Set(bit uint64) {
	var byteNum = bit / 8
	var bitNum = bit % 8
	if uint64(len(facet.bits)) < byteNum+1 {
		newBits := make([]byte, byteNum*2+1) // reserve double size for future grouth
		copy(newBits, facet.bits)
		facet.bits = newBits
	}
	facet.bits[byteNum] = facet.bits[byteNum] | (1 << bitNum)
}

func (facet *Facet) Count() (count uint64) {
	for _, bits := range facet.bits {
		for i := 0; i < 8 && bits > 0; i++ {
			if bits&1 == 1 {
				count++
			}
			bits >>= 1
		}
	}
	return
}
