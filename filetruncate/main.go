package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		panic("Expected input and output file names")
	}

	fin, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	fout, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	data := make([]byte, 1024*1024)
	foundNonZero := false
	zeroes := 0
	for {
		n, err := fin.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		data = data[:n]

		nNonZero := 0
		if !foundNonZero {
			for i, b := range data {
				if b != 0 {
					nNonZero = i
					foundNonZero = true
					zeroes += nNonZero
					break
				}
			}
		}

		if foundNonZero {
			fout.Write(data[nNonZero:])
			fmt.Print(".")
		} else {
			fmt.Print("0")
		}
	}

	fmt.Println()
	fmt.Println("zeroes:", zeroes)
}
