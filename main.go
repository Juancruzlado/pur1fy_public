package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Feriado struct {
	Fecha  string `json:"fecha"`
	Nombre string `json:"nombre"`
}

func main() {
	fmt.Println("[+] Initial log!\n")
	fmt.Println("[+] Scanning info...\n")

	res, err := http.Get("https://api.argentinadatos.com/v1/feriados/2024")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("Resource not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var feriados []Feriado
	err = json.Unmarshal(body, &feriados)
	if err != nil {
		panic(err)
	}

	for _, feriado := range feriados {
		fmt.Printf("Date: %s, Name: %s\n", feriado.Fecha, feriado.Nombre)
	}

	fmt.Println("\n")
	fmt.Println("[+] Scan finished as expected, bye!\n")
}
