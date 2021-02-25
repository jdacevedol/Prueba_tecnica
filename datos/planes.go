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

type Planes struct{

		codigo				primitive.ObjectID 	`yml:"_id,omitempty" bson:"_id,imutempty"`
		nombrePlan			string 				`yml:"nombrePlan,omitempty" bsom:"nombrePlan,omitempty"`
		aseguradora 		string				`yml:"aseguradora,omitempty" bsom:"aseguradora,omitempty"`
		montoPesos			float32				`yml:"montoPesos,omitempty" bsom:"montoPesos,omitempty"`
		montoUf				float32				`yml:"montoUf,omitempty" bsom:"montoUf,omitempty"`
		capitalAsegurado	float32				`yml:"capitalAsegurado,omitempty" bsom:"capitalAsegurado,omitempty"`
		valorPromocion		float32				`yml:"valorPromocion,omitempty" bsom:"valorPromocion,omitempty"`

}


func ObtenerPlanEndpoint(response http.ResponseWriter, request *http.Request) {

		response.Header().Set("content-type", "application/yml")


		parametros := mux.Vars(request)
		id, _ := primitive.ObjectIDFromHex(params["id"])
		var planes Planes


		coleccion := client.Database("pruebaTecnica").Collection("usuario")
		contexto, _ := context.WithTimeout(context.Background(), 30*time.Second)
		
		
		err := collection.FindOne(ctx, Planes{ID: id}).Decode(&planes)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}


		yml.NewEncoder(response).Encode(person)

}