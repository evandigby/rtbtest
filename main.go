// gotest project main.go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Required arguments missing (expecting port and file)")
		return
	}

	data, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}

	result, err := http.Post("http://localhost:"+os.Args[1], "application/json", bytes.NewReader(data))

	if err != nil {
		panic(err)
	}

	fmt.Println(os.Args)
	fmt.Println(result.Status)

	body, err := ioutil.ReadAll(result.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body[:]))
}
