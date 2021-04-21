package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"strings"
)

var nSize int

func GetRandomBytes(size int) ([]byte, error) {
	b := make([]byte, size)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func rawBytesToString(b []byte) string {
	var s = ""

	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%d ", b[i])
	}

	return strings.TrimSpace(s)
}

type point struct {
	x, y int
}

func main() {
	flag.IntVar(&nSize, "size", 32, "byte size")

	b, err := GetRandomBytes(nSize)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Byte Size: %d\n", nSize)

	rawEnc := base64.RawURLEncoding.EncodeToString(b)
	fmt.Printf("Raw: %s\n", rawEnc)

	urlEnc := base64.URLEncoding.EncodeToString(b)
	fmt.Printf("Url: %s\n", urlEnc)

	rawBytes, err := base64.RawURLEncoding.DecodeString(rawEnc)
	if err != nil {
		return
	}

	sBytes := rawBytesToString(rawBytes)
	fmt.Printf("Raw Bytes: %s\n", sBytes)

	fmt.Printf("|%6s|%6s|\n", "a", "b")
	fmt.Printf("|%-6s|%-6s|\n", "a", "b")

	points := []string{"a", "b"}
	fmt.Printf("len(points)\t%d\n", len(points))

	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%#v\n", p)
	p.x = 3
	fmt.Printf("%+v\n", p)
	fmt.Printf("%T\n", p)

}
