package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	r "math/rand"
)

func main() {

	l := Line{
		B:    &B{b: 1},
		Data: []byte("hello world"),
	}

	for i := 0; i < 50; i++ {
		d := r.Int31() % 127
		l.Data = append(l.Data, byte(d))
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("r: %d\tc: %d\n", r.Int31(), crand())
	}

	// fmt.Println("n:", n)

	l.Print()
}

func crand() *big.Int {
	n, err := rand.Int(rand.Reader, big.NewInt(1234567890))
	if err != nil {
		fmt.Println(err)
		return big.NewInt(0)
	}

	return n
}

type B struct{ b int32 }

type Line struct {
	B    *B
	Data []byte
}

func (l *Line) String() string {
	return fmt.Sprintf("%s%+v", l.B, l.Data)
}

func (b *B) String() string {
	return fmt.Sprintf("%d-", b.b)
}

func (l *Line) Print() {
	fmt.Printf("%s", l)
}
