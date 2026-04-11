package main

import (
	"fmt"

	"context"
	"log"

	"github.com/goprojects/hotel-reservation/db"
	"github.com/goprojects/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
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

	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)

	hotel := types.Hotel{
		Name:     "Kamare Valley",
		Location: "Palghar",
	}

	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 999,
		},
		{
			Type:      types.DoubleRoomType,
			BasePrice: 1499,
		},
		{
			Type:      types.SeasideRoomType,
			BasePrice: 1999,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for i := range rooms {
		rooms[i].HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &rooms[i])
		if err != nil {
			log.Fatal(err)
		} else {
			err := hotelStore.Update(ctx, bson.M{}, bson.M{})
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println(insertedRoom)
	}

	fmt.Println(insertedHotel)
}
