package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func main() {
	log.SetOutput(ioutil.Discard) // Switch off logging

	ctext, err := hex.DecodeString("9F0B13944841A832B2421B9EAF6D9836813EC9D944A5C8347A7CA69AA34D8DC0DF70E343C4000A2AE35874CE75E64C31")
	assert(err, "Error in hex decoding")

	o, err := NewOracle(16)
	assert(err, "Could not create oracle:")
	defer o.Close()

	r, err := o.Send(ctext)
	fmt.Printf("For initial code Oracle returned %c\n", r)
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
func (o *Oracle) Send(ctext []byte) (r byte, err error) {
	if len(ctext)%o.blocklen != 0 {
		return 0, fmt.Errorf("Message %q is not a whole multiple of block length %v", ctext, o.blocklen)
	}

	nblocks := len(ctext) / o.blocklen

	_, err = o.write([]byte{byte(nblocks)}, ctext, []byte{0})
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
