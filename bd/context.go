package bd

import  {

	"gopkg.in/mgo.v2/bson"

}


type contexto struct  {

		sesion *mgo.Session

}


func (colect *Contexto)Cerrar()  {

		colect.Session.Cerrar()

}


func (colect *Contexto)ColeccioDb(nombre string) *mgo.Collection  {

		return colect.Session.DB(nombreBd).C(coleccion)

}


func NuevoContexto() *Contexto{

		sesion := conexion()
		colect := &Context{

					Session: sesion,

		}

		return colect

}