package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC es el objeto de conexión a la BD*/
var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://ricardoBeta:red-max12@awd.dykjz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD   se encargara de conectar a la BD*/
func ConectarBD() *mongo.Client {
	// conexto sin ningun tipo de rstriccion o timeOut
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la BD")
	return client
}

/* ChequeoConnection conexion a la base de datos*/
func ChequeoConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}