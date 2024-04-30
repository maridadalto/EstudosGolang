package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	ID        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Endereco  *Endereco `json:"endereco"`
}

type Endereco struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var Pessoas []Pessoa = []Pessoa{
	Pessoa{
		ID:        "1",
		Firstname: "Maria",
		Lastname:  "dos Santos",
		Endereco: &Endereco{
			City:  "Sorocaba",
			State: "SP",
		},
	},
	Pessoa{
		ID:        "2",
		Firstname: "Paulo",
		Lastname:  "Ribeiro",
		Endereco: &Endereco{
			City:  "Sorocaba",
			State: "SP",
		},
	},
}

// Lista todos os contatos
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Pessoas)
}

// Retorna apenas o contato selecionado pelo ID
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range Pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pessoa{})
}

// Cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var pessoa Pessoa
	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoa.ID = params["id"]
	Pessoas = append(Pessoas, pessoa)
	json.NewEncoder(w).Encode(Pessoas)
}

// Deleta contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range Pessoas {
		if item.ID == params["id"] {
			Pessoas = append(Pessoas[:index], Pessoas[index+1:]...)
			break

		}
		json.NewEncoder(w).Encode(Pessoas)
	}
}

func Rotas(roteador *mux.Router) {
	roteador.HandleFunc("/contato", GetPeople).Methods("GET")
	roteador.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	roteador.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	roteador.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
}
func jsonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//transforma em json
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func configurarServidor() {
	roteador := mux.NewRouter().StrictSlash(true)
	roteador.Use(jsonMiddleWare)
	Rotas(roteador)
	//rodar o servidor
	log.Fatal(http.ListenAndServe(":8080", roteador))
}

func main() {
	configurarServidor()
}
