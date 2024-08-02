package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Initial log!")
	res, err := http.Get("https://api.argentinadatos.com/v1/feriados/2024")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Resource not avaliable")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
