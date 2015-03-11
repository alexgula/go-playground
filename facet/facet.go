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
	var byteNum, bitNum = f.pos(bit)
	if uint(len(f.bits)) <= byteNum {
		newBits := make([]byte, byteNum*2+1) // reserve double size for future grouth
		copy(newBits, f.bits)
		f.bits = newBits
	}
	f.bits[byteNum] = f.bits[byteNum] | (1 << bitNum)
}

func (f *Facet) Clear(bit uint) {
	var byteNum, bitNum = f.pos(bit)
	if uint(len(f.bits)) <= byteNum {
		return // Non-existent bits are considered cleared
	}
	f.bits[byteNum] = f.bits[byteNum] & ^(1 << bitNum)
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

func (f *Facet) And(o *Facet) {
	if len(f.bits) > len(o.bits) {
		f.bits = f.bits[:len(o.bits)]
	}
	for i := 0; i < len(f.bits); i++ {
		f.bits[i] &= o.bits[i]
	}
}

func (f *Facet) pos(bit uint) (byteNum uint, bitNum uint) {
	byteNum = bit / itemsize
	bitNum = bit % itemsize
	return
}
