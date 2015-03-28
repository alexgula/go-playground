package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type codeStats [256]float64

func main() {
	log.SetOutput(ioutil.Discard) // Switch off logging

	c, err := hex.DecodeString("C0DE")
	assert(err, "Error in hex decoding")

	n, err := findKeyLen(c, 13)
	assert(err, "Error in key length")

	fmt.Printf("N = %d\n", n)

	k, err := findKey(c, n)
	assert(err, "Error in find key")

	fmt.Printf("Key = %v\n", k)

	t, err := bytesToSlice(xor(bytes.NewReader(c), k...))
	assert(err, "Error in applying key")

	fmt.Println(string(t))
}

func assert(err error, msg string) {
	if err != nil {
		panic(fmt.Sprint(msg, err))
	}
}

func findKeyLen(c []byte, maxLen int) (int, error) {
	var maxi int
	var maxd float64
	for i := 1; i <= maxLen; i++ {
		logBuf := bytes.NewBufferString("")
		fmt.Fprintf(logBuf, "%2d", i)
		d := float64(0)
		for j := 0; j < i; j++ {
			stats, err := newCodeStats(filter(bytes.NewReader(c), j, i))
			if err != nil {
				return 0, err
			}
			dj := stats.d()
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

func findKey(c []byte, n int) ([]byte, error) {
	k := make([]byte, n)
	for i := 0; i < n; i++ {
		var maxj byte
		var maxd float64
		for j := 0; j < 256; j++ {
			stats, err := newCodeStats(xor(filter(bytes.NewReader(c), i, n), byte(j)))
			if err != nil {
				return nil, err
			}
			d := mulCodeStats(stats, newEnglishStats())
			if d > maxd {
				maxj, maxd = byte(j), d
			}
		}
		k[i] = maxj
	}
	return k, nil
}

func newCodeStats(r io.ByteReader) (codeStats, error) {
	s := codeStats{}

	var n int
	for {
		b, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				return codeStats{}, err
			}
			break
		}
		s[b]++
		n++
	}

	if n == 0 {
		return s, nil
	}

	for i := 0; i < 256; i++ {
		s[i] /= float64(n)
	}

	return s, nil
}

func newEnglishStats() codeStats {
	e := codeStats{
		'a': 8.167,
		'b': 1.492,
		'c': 2.782,
		'd': 4.253,
		'e': 12.702,
		'f': 2.228,
		'g': 2.015,
		'h': 6.094,
		'i': 6.966,
		'j': 0.153,
		'k': 0.772,
		'l': 4.025,
		'm': 2.406,
		'n': 6.749,
		'o': 7.507,
		'p': 1.929,
		'q': 0.095,
		'r': 5.987,
		's': 6.327,
		't': 9.056,
		'u': 2.758,
		'v': 0.978,
		'w': 2.360,
		'x': 0.150,
		'y': 1.974,
		'z': 0.074,
	}
	for i := 0; i < len(e); i++ {
		e[i] /= 100
	}
	return e
}

func (s codeStats) d() float64 {
	return mulCodeStats(s, s)
}

func mulCodeStats(s1, s2 codeStats) float64 {
	var m float64
	for i := 0; i < 256; i++ {
		m += s1[i] * s2[i]
	}
	return m
}

type periodFilter struct {
	src           io.ByteReader
	start, period int
}

func (f *periodFilter) ReadByte() (c byte, err error) {
	for i := 0; i <= f.start; i++ {
		c, err = f.src.ReadByte()
		if err != nil {
			return
		}
	}
	f.start = 0 // Skip first bytes on first run only

	for i := 1; i < f.period; i++ {
		_, err = f.src.ReadByte()
		if err != nil {
			return
		}
	}
	return
}

func filter(src io.ByteReader, start, period int) io.ByteReader {
	return &periodFilter{src, start, period}
}

type xorStream struct {
	src io.ByteReader
	pos int
	key []byte
}

func xor(src io.ByteReader, key ...byte) io.ByteReader {
	return &xorStream{src, 0, key}
}

func (s *xorStream) ReadByte() (c byte, err error) {
	c, err = s.src.ReadByte()
	if err != nil {
		return
	}
	c ^= s.key[s.pos%len(s.key)]
	s.pos++
	return
}

func bytesToSlice(r io.ByteReader) (b []byte, err error) {
	for {
		c, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				return b, err
			}
			break
		}
		b = append(b, c)
	}
	return
}
