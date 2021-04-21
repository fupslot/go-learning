package main

import (
	"encoding/base64"
	"fmt"
	"time"
)

func main() {

	str := "привет world"

	b64EStr := base64.StdEncoding.EncodeToString([]byte(str))
	b64UStr := base64.URLEncoding.EncodeToString([]byte(str))

	fmt.Println(b64EStr)
	fmt.Println(b64UStr)

	ch := make(chan string)

	if v, err := base64.StdEncoding.DecodeString(b64EStr); err == nil {
		fmt.Println(string(v))
	}

	go func(bs []byte, ch chan string) {
		for _, v := range bs {
			time.Sleep(5 * time.Second)
			ch <- string(v)
		}

		close(ch)
	}([]byte{100, 87, 54}, ch)

	for out := range ch {
		fmt.Println(out)
	}

	fmt.Println("Done!")
}
