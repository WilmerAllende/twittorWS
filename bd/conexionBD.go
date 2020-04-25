package bd

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

)

var MongoCN = conectarBD() 
var clientOptions = options.Client().ApplyURI("mongodb+srv://wsanchez:uzumaki2010@cluster0-uutdn.mongodb.net/test?retryWrites=true&w=majority")

func conectarBD() *mongo.Client  {
    client,error := mongo.Connect(context.TODO(),clientOptions)
    if error != nil {
        log.Fatal(error.Error())
        return client
    }

    error = client.Ping(context.TODO(),nil)
    if error != nil {
        log.Fatal(error.Error())
        return client
    }

    log.Println("Conexion exitosa con BD")
    return client
}

func ChequeoConnection() int {
    error := MongoCN.Ping(context.TODO(),nil)
    if error != nil {
        return 0
    }else{
        return 1;
    }
}