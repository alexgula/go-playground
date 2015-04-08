package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexgula/go-playground/crypto/filter"
	"github.com/alexgula/go-playground/crypto/stats"
	"github.com/alexgula/go-playground/crypto/xor"
)

func main() {
	log.SetOutput(ioutil.Discard) // Switch off logging

	buf := make([]byte, 1024)

	c, err := hex.DecodeString("C0DE")
	assert(err, "Error in hex decoding")

	n, err := findKeyLen(bytes.NewReader(c), 13, buf)
	assert(err, "Error in key length")

	fmt.Printf("N = %d\n", n)

	k, err := findKey(bytes.NewReader(c), n, buf)
	assert(err, "Error in find key")

	fmt.Printf("Key = %v\n", k)

	_, err = printR(xor.NewReader(bytes.NewReader(c), k...))
	assert(err, "Error in applying key")
}

func assert(err error, msg string) {
	if err != nil {
		panic(fmt.Sprint(msg, err))
	}
}

func printR(src io.Reader) (n int64, err error) {
	n, err = io.Copy(os.Stdout, src)
	fmt.Println()
	return
}

func findKeyLen(src io.ReadSeeker, maxLen int, buf []byte) (int, error) {
	var maxi int
	var maxd float64
	for i := 1; i <= maxLen; i++ {
		logBuf := bytes.NewBufferString("")
		fmt.Fprintf(logBuf, "%2d", i)
		d := float64(0)
		for j := 0; j < i; j++ {
			src.Seek(0, 0)
			s, err := stats.NewFrom(filter.NewReader(src, j, i), buf)
			if err != nil {
				return 0, err
			}
			dj := s.D()
			fmt.Fprintf(logBuf, " %6.2f", dj*100)
			d += dj / float64(i)
		}
		if d > maxd {
			maxi, maxd = i, d
		}
		fmt.Fprintf(logBuf, "  ->  %6.2f", d*100)
		log.Println(logBuf)
	}
	return maxi, nil
}

func findKey(src io.ReadSeeker, keyLen int, buf []byte) ([]byte, error) {
	k := make([]byte, keyLen)
	for i := 0; i < keyLen; i++ {
		var maxj byte
		var maxd float64
		for j := 0; j < 256; j++ {
			src.Seek(0, 0)
			s, err := stats.NewFrom(xor.NewReader(filter.NewReader(src, i, keyLen), byte(j)), buf)
			if err != nil {
				return nil, err
			}
			d := stats.Mul(s, stats.English())
			if d > maxd {
				maxj, maxd = byte(j), d
			}
		}
		k[i] = maxj
	}
	return k, nil
}
