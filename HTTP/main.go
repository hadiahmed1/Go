package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Printf("%v",resp)
	bs := make([]byte, 999)
	resp.Body.Read((bs)) // pushing body in ByteSlice

	body := string(bs)
	fmt.Println(body)

	io.Copy(os.Stdout, resp.Body)
}
