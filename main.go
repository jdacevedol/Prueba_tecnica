package main


import (
		"context"
		"encoding/yml"
		"fmt"
		"net/http"
		"time"

		"github.com/gorilla/mux"
		"go.mongodb.org/mongo-driver/bson/primitive"
		"go.mongodb.org/mongo-driver/mongo"
)


var cliente *mongo.Cliente


func Conexion() {

		fmt.Println("Iniciando aplicacion")

		contexto, _ := context.WithTimeout(context.Background(), 10*time.Second)
		opcionesCliente := options.Cliente().ApplyURI("mongodb://localhost:27017")
		cliente, _ = mongo.Connect(contexto, opcionesCliente)

}


func main() {

		conexion()

		router := mux.NewRouter()
		router.HandleFunc("/usuario", CrearUsuarioEndpoint).Methods("POST")
		router.HandleFunc("/planes", GetPeopleEndpoint).Methods("GET")


		http.ListenAndServe(":12345", router)

}
