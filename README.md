**DOCKER PROOF OF CONCEPT**

This repository contains a simple REST test program that has been dockerized.

The test program is written Go and is shown below:
~~~~~~~~~~~~~~~~~~~~~~~~~GOLANG
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
~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Included is **docker.md.html** detailing how to setup docker images in general and how the, **dockgo** image, included, was built.

