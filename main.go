package main

import (
	"context"
	_ "fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/goprojects/hotel-reservation/api"
	"github.com/goprojects/hotel-reservation/db"
	_ "github.com/goprojects/hotel-reservation/types"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017/"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	//Learning Code
	/*	ctx := context.Background()
		coll := client.Database(dbname).Collection(userColl)

		user := types.User{
			FirstName: "Siddhant",
			LastName:  "Tamboli",
		}
		_, err = coll.InsertOne(ctx, user)
		if err != nil {
			log.Fatal(err)
		}

		var bachua types.User
		if err = coll.FindOne(ctx, bson.M{}).Decode(&bachua); err != nil {
			log.Fatal(err)
		}
		fmt.Println(bachua)*/

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DBNAME))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user:id", userHandler.HandleDeleteUser)
	app.Listen(":4040")
}
