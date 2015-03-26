package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

type codeStats [256]float64

func main() {
	log.SetOutput(ioutil.Discard) // Switch off logging

	c, err := hex.DecodeString("C0DE")
	if err != nil {
		fmt.Println("Error in hex decoding", err)
	}

	n := findKeyLen(c, 13)

	fmt.Printf("N = %d\n", n)

	k := findKey(c, n)

	fmt.Printf("Key = %v\n", k)

	t := xor(c, k...)

	fmt.Println(string(t))
}

func findKeyLen(c []byte, maxLen int) int {
	maxi, maxd := 0, float64(0)
	for i := 1; i <= maxLen; i++ {
		logBuf := bytes.NewBufferString("")
		fmt.Fprintf(logBuf, "%2d", i)
		d := float64(0)
		for j := 0; j < i; j++ {
			dj := newCodeStats(filter(c, j, i)).d()
			fmt.Fprintf(logBuf, " %6.2f", dj*100)
			d += dj / float64(i)
		}
		if d > maxd {
			maxi, maxd = i, d
		}
		fmt.Fprintf(logBuf, "  ->  %6.2f", d*100)
		log.Println(logBuf)
	}
	return maxi
}

func findKey(c []byte, n int) []byte {
	k := make([]byte, n)
	for i := 0; i < n; i++ {
		maxj, maxd := byte(0), float64(0)
		for j := 0; j < 256; j++ {
			d := mulCodeStats(newCodeStats(xor(filter(c, i, n), byte(j))), newEnglishStats())
			if d > maxd {
				maxj, maxd = byte(j), d
			}
		}
		k[i] = maxj
	}
	return k
}

func newCodeStats(c []byte) codeStats {
	s := codeStats{}
	f := 1 / float64(len(c))
	for i := 0; i < len(c); i++ {
		s[c[i]] += f
	}
	return s
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
	return e
}

func (s codeStats) d() float64 {
	return mulCodeStats(s, s)
}

func mulCodeStats(s1, s2 codeStats) float64 {
	var m float64 = 0.0
	for i := 0; i < 256; i++ {
		m += s1[i] * s2[i]
	}
	return m
}

func avg(s []float64) float64 {
	var r float64 = 0.0
	for _, v := range s {
		r += v
	}
	return r / float64(len(s))
}

func filter(s []byte, t, p int) []byte {
	r := make([]byte, int(math.Ceil(float64(len(s)-t)/float64(p))))
	for i, j := 0, t; j < len(s); i, j = i+1, j+p {
		r[i] = s[j]
	}
	return r
}

func xor(s []byte, k ...byte) []byte {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[i] ^ k[i%len(k)]
	}
	return r
}
