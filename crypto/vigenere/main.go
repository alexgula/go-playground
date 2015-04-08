package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type codeStats [256]int

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

	_, err = printR(xor(bytes.NewReader(c), k...))
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
			stats, err := newCodeStats(filter(src, j, i), buf)
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

func findKey(src io.ReadSeeker, keyLen int, buf []byte) ([]byte, error) {
	k := make([]byte, keyLen)
	for i := 0; i < keyLen; i++ {
		var maxj byte
		var maxd float64
		for j := 0; j < 256; j++ {
			src.Seek(0, 0)
			stats, err := newCodeStats(xor(filter(src, i, keyLen), byte(j)), buf)
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

func newCodeStats(src io.Reader, buf []byte) (s codeStats, err error) {
	var n int
	for err == nil {
		n, err = src.Read(buf)
		s.add(buf[:n])
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func (s *codeStats) add(p []byte) {
	for i := 0; i < len(p); i++ {
		s[p[i]]++
	}
}

func newEnglishStats() codeStats {
	e := codeStats{
		'a': 8167,
		'b': 1492,
		'c': 2782,
		'd': 4253,
		'e': 12702,
		'f': 2228,
		'g': 2015,
		'h': 6094,
		'i': 6966,
		'j': 153,
		'k': 772,
		'l': 4025,
		'm': 2406,
		'n': 6749,
		'o': 7507,
		'p': 1929,
		'q': 95,
		'r': 5987,
		's': 6327,
		't': 9056,
		'u': 2758,
		'v': 978,
		'w': 2360,
		'x': 150,
		'y': 1974,
		'z': 74,
	}
	return e
}

func (s codeStats) d() float64 {
	return mulCodeStats(s, s)
}

func (s codeStats) total() int {
	t := 0
	for _, c := range s {
		t += c
	}
	return t
}

func mulCodeStats(s1, s2 codeStats) float64 {
	t1, t2 := s1.total(), s2.total()
	if t1 == 0 || t2 == 0 {
		return 0
	}
	var m float64
	for i := 0; i < 256; i++ {
		m += (float64(s1[i]) / float64(t1)) * (float64(s2[i]) / float64(t2))
	}
	return m
}

type periodFilter struct {
	src           io.Reader
	pos           int
	start, period int
}

func (f *periodFilter) Read(p []byte) (n int, err error) {
	n, err = f.src.Read(p)
	n = f.filter(p[:n])
	return
}

func (f *periodFilter) filter(p []byte) (n int) {
	for i := f.start; i < len(p); i, n = i+f.period, n+1 {
		p[n] = p[i]
	}

	// Skip first bytes on first run only
	f.start -= len(p)
	if f.start < 0 {
		f.start = 0
	}

	return
}

func filter(src io.Reader, start, period int) io.Reader {
	return &periodFilter{src: src, start: start, period: period}
}

type xorStream struct {
	src io.Reader
	pos int
	key []byte
}

func xor(src io.Reader, key ...byte) io.Reader {
	return &xorStream{src: src, key: key}
}

func (s *xorStream) Read(p []byte) (n int, err error) {
	n, err = s.src.Read(p)
	s.xor(p[:n])
	return
}

func (s *xorStream) xor(p []byte) {
	if len(s.key) == 0 {
		return
	}
	for i := 0; i < len(p); i, s.pos = i+1, s.pos+1 {
		p[i] ^= s.key[s.pos%len(s.key)]
	}
}
