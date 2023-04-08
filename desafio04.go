package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type clientes struct {
	Nome       string `json:"Nome"`
	Sabor      string `json:"Sabor"`
	Confeitado string `json:"Confeitado"`
	Valor      string `json:"Valor"`
}
type resposta struct {
	Clientes []clientes `json:"Clientes"`
}

func main() {
	file, err := ioutil.ReadFile("apiRest.json")
	if err != nil {
		log.Fatal(err)
	}

	var resp resposta
	if err := json.Unmarshal(file, &resp); err != nil {
		log.Fatal(err)
	}

	for i := range resp.Clientes {
		valor, err := strconv.ParseFloat(fmt.Sprintf("%v", resp.Clientes[i].Valor), 64)
		if err != nil {
			log.Fatal(err)
		}
		resp.Clientes[i].Valor = fmt.Sprintf("%.2f", valor)
	}

	r := mux.NewRouter()

	r.HandleFunc("/bolos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp.Clientes); err != nil {
			log.Fatal(err)
		}
	})

	fmt.Println("Servidor rodando em http://localhost:8080") //http://localhost:8080/bolos//
	log.Fatal(http.ListenAndServe(":8080", r))
}
