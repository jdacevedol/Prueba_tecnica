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


type Usuario struct {

	ID              primitive.ObjectID `yml:"_id,omitempty" bson:"_id,omitempty"`
	ruta            int                `yml:"ruta,omitempty" bson:"ruta,omitempty"`
	fechaNacimiento string             `yml:"fechaNacimiento,omitempty" bson:"fechaNacimiento,omitempty"`
	correo          string             `yml:"correo,omitempty" bson:"correo,omitempty"`
	telefono        int                `yml:"telefono,omitempty" bson:"telefono,omitempty"`
	
}


func CrearUsuarioEndpoint(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/yml")


	var usuario Usuario
	_ = yml.NewDecoder(request.Body).Decode(&usuario)


	coleccion := cliente.Database("pruebaTecnica").Collection("usuario")
	contexto, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := coleccion.InsertOne(contexto, usuario)


	yml.NewEncoder(response).Encode(result)

}
