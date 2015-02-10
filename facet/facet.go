package facet

type Facet struct {
	bits []byte
}

func New() *Facet {
	return &Facet{}
}

func (f *Facet) Set(bit uint) {
	var byteNum = bit / 8
	var bitNum = bit % 8
	if uint(len(f.bits)) < byteNum+1 {
		newBits := make([]byte, byteNum*2+1) // reserve double size for future grouth
		copy(newBits, f.bits)
		f.bits = newBits
	}
	f.bits[byteNum] = f.bits[byteNum] | (1 << bitNum)
}

func (f *Facet) Count() uint {
	var n uint = 0
	for _, bits := range f.bits {
		for i := 0; i < 8 && bits > 0; i++ {
			if bits&1 == 1 {
				n++
			}
			bits >>= 1
		}
	}
	return n
}
