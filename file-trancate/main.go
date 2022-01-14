package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./index", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error())
	}

	size := fileSize(f)

	fmt.Printf("File size (before trancate) %d\n", size)

	os.Truncate(f.Name(), int64(32651))

	fmt.Printf("File size (after trancate) %d\n", fileSize(f))

	n := writeFile(f, size, []byte("hello world"))
	os.Truncate(f.Name(), int64(size+n))

	fmt.Printf("File size (after trancate) %d\n", fileSize(f))
}

func writeFile(f *os.File, pos int64, b []byte) int64 {
	n, err := f.WriteAt(b, pos)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return int64(n)
}

func fileSize(f *os.File) int64 {
	fi, err := os.Stat(f.Name())
	if err != nil {
		panic(err.Error())
	}

	return fi.Size()
}
