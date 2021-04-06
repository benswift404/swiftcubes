package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
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
	// Set cookie and execute template
	id := uuid.Must(uuid.NewRandom())
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: id.String(),
	})
	tpl.ExecuteTemplate(w, "index.gohtml", cubes)
}

func cubeDetail(w http.ResponseWriter, r *http.Request) {
	// Get cookie or return an error if there is no cookie found
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Fprintf(w, "No cookie")
		return
	}

	// Get URL variable, convert to int, then put cube instance into cube variable
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	cube := cubes[id_int]

	// Get data packed into struct for the template
	data := struct {
		Cube   Cube
		Cookie *http.Cookie
	}{cube, c}
	tpl.ExecuteTemplate(w, "detail.gohtml", data)
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
