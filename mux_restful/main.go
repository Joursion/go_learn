package main

import (
  "log"
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
)

type Person struct {
  ID string `json:"id,omitempty"` // 为空的时候忽略
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Address *Address `json:"address,omitempty"`
}

type Address struct {
  City string `json:"city,omitempyt"`
  Provice string `json:"province,omitempty"`
}

var people []Person

func GetPerson (w http.ResponseWriter, req * http.Request) {
  params := mux.Vars(req)
  for _, item := range(people) {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(people)
}

func GetPeople (w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(people)
}

func PostPerson (w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  var person Person
  _ = json.NewDecoder(req.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}

func DelPerson (w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for index, item := range people {
    if item.ID == params["id"] {
      people = append(people[:index], people[index + 1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(people)
}

func main () {
  people = append(people, Person{ID: "1", Firstname: "1", Lastname: "1111", Address: &Address{City: "1111111", Provice: "11111"}})  
  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", PostPerson).Methods("POST")
  router.HandleFunc("/people/{id}", DelPerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":12345", router))
}