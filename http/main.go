package main

import (
	"fmt"
	"mime"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// defer res.Body.Close()

	res.Header.Del("Set-Cookie")
	mediatype, params, err := mime.ParseMediaType(res.Header.Get("Content-Type"))

	fmt.Printf("\t%s\n", mediatype)
	fmt.Printf("\t%s\n", params["charset"])
	fmt.Printf("\t%s\n", res.Header.Get("Content-Type"))
	fmt.Printf("\t%s\n", mime.TypeByExtension("xml"))
	fmt.Println("----")

	printHeaders(res.Header)

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%s", body)
}

func printHeaders(h http.Header) {
	for k, v := range h {
		fmt.Printf("%s\t%s\n", k, v)
	}
}
