package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type logWriter struct{}

func main() {
	r := strings.NewReader("Hello world!\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println("Error", err)
	}

	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	// fmt.Println(string(bs))
	return os.Stdout.Write(bs)
	// return len(bs), nil
}
