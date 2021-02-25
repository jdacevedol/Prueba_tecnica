package usuario


import  {

		"gopkg.in/mgo.v2/bson"

}


type Alamacenamiento interface  {
	 

		create(usuario *Usuario) error

}


var alamacenamiento Alamacenamiento


func cambiarAlmacenamiento(guardar Alamacenamiento){

		alamacenamiento = guardar

}


type Usuario struct  {


		ID bson.Objectid 	`bson:"_id,omitempty"`
		ruta				int
		fechaNacimiento		string
		correo				string
		telefono			int

}