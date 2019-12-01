package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main3() {
	resp, _ := http.Get("http://example.com/")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
