package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Call Get()
	resp, err := http.Get("https://naver.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Print Result
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
}
