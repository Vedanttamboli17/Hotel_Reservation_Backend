package main

import (
	"fmt"

	"github.com/goprojects/hotel-reservation/db"
	"github.com/goprojects/hotel-reservation/types"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name: "Sai Recidency",
		Location: "Palghar",
	}

	room := types.Room{
		Type: types.SingleRoomType,
		BasePrice: 1999,
	}
	_ = room
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(insertedHotel)
}
