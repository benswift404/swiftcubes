package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Cube struct {
	Id       int
	Name     string
	Category string
	Price    int
}

var tpl *template.Template
var cubes []Cube

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", cubes)
}

func cubeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	cube := cubes[id_int]
	tpl.ExecuteTemplate(w, "detail.gohtml", cube)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	cubes = append(cubes, Cube{
		Id:       0,
		Name:     "DaYan Tengyun V2 M",
		Category: "3x3",
		Price:    20,
	})

	cubes = append(cubes, Cube{
		Id:       1,
		Name:     "QiYi Megamorphix",
		Category: "Shape Mod",
		Price:    6,
	})

	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/cubes/{id}", cubeDetail).Methods("GET")

	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
