package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//log.SetOutput(ioutil.Discard) // Switch off logging

	ctext, err := hex.DecodeString("9F0B13944841A832B2421B9EAF6D9836813EC9D944A5C8347A7CA69AA34D8DC0DF70E343C4000A2AE35874CE75E64C31")
	assert(err, "Error in hex decoding")

	o, err := NewOracle(16)
	assert(err, "Could not create oracle:")
	defer o.Close()

	r, err := o.Send(ctext, -1, 0)
	log.Printf("For initial code oracle returned %c\n", r)

	cn := len(ctext) - o.blocklen
	padding := 0
	for i := cn - o.blocklen; i < cn; i++ {
		r1, err := o.Send(ctext, i, 0)
		assert(err, "Error checking substitution by 0")
		r2, err := o.Send(ctext, i, 1)
		assert(err, "Error checking substitution by 1")
		log.Printf("For i=%2d oracle returned %c and %c\n", i, r1, r2)
		if r1 == '1' && r2 == '1' {
			padding = cn - i - 1
		}
	}
	log.Printf("Padding is %d", padding)

	message := make([]byte, cn-padding)
	f := make([]byte, cn)
	for i := cn - padding; i < cn; i++ {
		f[i] = ctext[i] ^ byte(padding)
	}

	for i := cn - padding - 1; i >= 0; i-- {
		padding++
		if padding > o.blocklen {
			padding = 1
			cn -= o.blocklen
		}
		tctext := make([]byte, cn+o.blocklen)
		copy(tctext, ctext)
		log.Printf("Padding = %d, i = %d, cn = %d, len(ctext) = %d", padding, i, cn, len(tctext))
		for j := i + 1; j < cn; j++ {
			tctext[j] = f[j] ^ byte(padding)
		}
		k := 0
		for ; k < 256; k++ {
			r, err := o.Send(tctext, i, byte(k))
			assert(err, "Error while trying to find k")
			if r == '1' {
				f[i] = byte(k) ^ byte(padding)
				break
			}
		}
		message[i] = f[i] ^ tctext[i]
		log.Printf("k = %d, f = %d, ctext = %d ---------------> message = %#U", k, f[i], tctext[i], message[i])
	}

	fmt.Printf("\n%s\n", message)
}

func assert(err error, msg string) {
	if err != nil {
		panic(fmt.Sprint(msg, err))
	}
}

// Oracle encapsulates network connection along with specified block length.
type Oracle struct {
	net.Conn
	blocklen int
}

// NewOracle returns new oracle with established connection and specified block
// length.
func NewOracle(blocklen int) (o *Oracle, err error) {
	conn, err := net.DialTimeout("tcp", "54.165.60.84:80", time.Second)
	if err != nil {
		return
	}
	o = &Oracle{
		Conn:     conn,
		blocklen: blocklen,
	}
	return
}

// Send writes specified text into commection, prepended by number of blocks
// and appended with zero byte (this is the requirement of external service).
// If message length is not multiple of block length, nothing is written.
func (o *Oracle) Send(ctext []byte, i int, subst byte) (r byte, err error) {
	if len(ctext)%o.blocklen != 0 {
		return 0, fmt.Errorf("Message %q is not a whole multiple of block length %v", ctext, o.blocklen)
	}

	nblocks := len(ctext) / o.blocklen

	if i < 0 || i >= len(ctext) {
		_, err = o.write([]byte{byte(nblocks)}, ctext, []byte{0})
	} else {
		_, err = o.write([]byte{byte(nblocks)}, ctext[:i], []byte{subst}, ctext[i+1:], []byte{0})
	}
	if err != nil {
		return
	}

	buf := make([]byte, 2)
	_, err = o.Read(buf)
	if err != nil {
		return
	}

	r = buf[0]
	return
}

func (o *Oracle) write(bs ...[]byte) (n int, err error) {
	var m int
	w := bufio.NewWriter(o.Conn)
	for _, b := range bs {
		m, err = w.Write(b)
		n += m
		if err != nil {
			return
		}
	}
	err = w.Flush()
	return
}
