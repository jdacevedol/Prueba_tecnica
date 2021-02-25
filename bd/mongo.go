package bd

cosnt  {

		nombreBd 	= "pruebaTecnica"
		coleccion 	= "usuario"

}


var sesion *mgo.Session
func conexion()  {

		var err error
		sesion, err = mgo.Dial(url:"localhost")
		if err != nil{
				log.Fatal(err)
		}

}


func estadoSesion() *mgo.Session  {

		if sesion == nil{
				conexion()
		}


		return sesion

}


func inicializar()  {

		conexion()

}