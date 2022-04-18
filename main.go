// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"time"

// 	//"github.com/fasthttp/router"
// 	"github.com/qiangxue/fasthttp-routing"
// 	"github.com/keploy/go-sdk/integrations/kmongo"
// 	"github.com/keploy/go-sdk/keploy"
// 	"github.com/valyala/fasthttp"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.uber.org/zap"
// )

// type Name struct {
// 	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitemepty"`
// 	FirstName string             `json:"firstname,omitempty" bson:"lastname,omitemepty"`
// 	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitemepty"`
// }

// func Index(ctx *routing.Context) error{
// 	ctx.WriteString("Welcome!")
// 	return nil

// }

// func Hello(ctx *routing.Context	) error{
// 	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
// 	return nil
// }

// func insertPerson(ctx *routing.Context) error {
// 	ctx.Response.Header.Add("content-type", "application/json")
// 	var person Name
// 	json.NewDecoder(ctx.RequestBodyStream()).Decode(&person)
// 	c, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	result, _ := col.InsertOne(c, person)
// 	json.NewEncoder(ctx.Response.BodyWriter()).Encode(result)
// 	return nil
// }

// // func insertPerson(n Name){
// // 	inserted, err:= col.InsertOne(context.Background(), n)
// // 	if err != nil{
// // 		logger.Fatal(err.Error())
// // 	}
// // 	fmt.Println("Inserted one person", inserted.InsertedID)
// // }

// var col *kmongo.Collection
// var logger *zap.Logger

// func New(host, db string) (*mongo.Client, error) {
// 	clientOptions := options.Client()

// 	clientOptions.ApplyURI("mongodb://" + host + "/" + db + "?retryWrites=true&w=majority")

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	return mongo.Connect(ctx, clientOptions)
// }

// func main() {
// 	logger, _ = zap.NewProduction()
// 	defer logger.Sync() // flushes buffer, if any

// 	dbName, collection := "info", "persons"
// 	client, err := New("localhost:27017", dbName)
// 	if err != nil {
// 		logger.Fatal("failed to create mgo db client", zap.Error(err))
// 	}
// 	db := client.Database(dbName)

// 	// integrate keploy with mongo
// 	col = kmongo.NewCollection(db.Collection(collection))

// 	r := routing.New()
// 	r.Get("/", Index)
// 	r.Get("/hello/{name}", Hello)
// 	r.Post("/person", insertPerson)
// 	k := keploy.New(keploy.Config{
// 		App: keploy.AppConfig{
// 			Name: "URL",
// 			Port: "8080",
// 		},
// 		Server: keploy.ServerConfig{
// 			URL: "http://localhost:8081/api",
// 		},
// 	})
// 	Fast(k,r)

// 	log.Fatal(fasthttp.ListenAndServe(":8080", r.HandleRequest))
// }
