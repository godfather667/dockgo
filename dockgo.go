package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("172.17.0.3000/movies")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(len(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
