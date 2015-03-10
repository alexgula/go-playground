package facet

type Facet struct {
	bits []byte
}

const (
	itemsize = 8
)

func New() *Facet {
	return &Facet{}
}

func (f *Facet) Set(bit uint) {
	var byteNum = bit / itemsize
	var bitNum = bit % itemsize
	if uint(len(f.bits)) <= byteNum {
		newBits := make([]byte, byteNum*2+1) // reserve double size for future grouth
		copy(newBits, f.bits)
		f.bits = newBits
	}
	f.bits[byteNum] = f.bits[byteNum] | (1 << bitNum)
}

func (f *Facet) Clear(bit uint) {
}

func (f *Facet) Count() uint {
	var n uint = 0
	for _, bits := range f.bits {
		for i := 0; i < itemsize && bits > 0; i++ {
			if bits&1 == 1 {
				n++
			}
			bits >>= 1
		}
	}
	return n
}
