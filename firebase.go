package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func database(data int) {
	// Get the database initalized
	client, ctx, err := FirebaseInstance()
	if err != nil {
		panic(err)
	}

	// Simple counter
	//var counter int

	for {
		// Example read of data
		fmt.Printf("data value %v\n", data)
		dsnap, err := client.Collection("orders").Doc("example").Get(ctx)
		if err != nil {
			panic(err)
		}
		var order Order
		dsnap.DataTo(&order)
		fmt.Printf("Read Document data: %#v\n", order)

		// Wait some period of time
		time.Sleep(time.Second * 2)

		//fmt.Printf("Increasing old Weight form %v to %v\n", order.Weight, counter)
		fmt.Printf("Update old Weight from %v to %v\n", order.Weight, data)

		// Update counter from the database
		order.Weight = data

		// Wait some period of time
		time.Sleep(time.Second * 2)

		// Example write of data
		_, err = client.Collection("orders").Doc("example").Update(ctx, []firestore.Update{
			{
				Path: "weight",
				//Value: counter,
				Value: data,
			},
		})
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			panic(err)
		} else {
			fmt.Println("Updated document in database")
		}
		// Wait some period of time
		time.Sleep(time.Second * 3)
	}
}

// FirebaseInstance obtains the client and ctx when needed
func FirebaseInstance() (*firestore.Client, context.Context, error) {

	// Get static variables for setting up the firestore
	var opt = option.WithCredentialsFile("private/smart-scale-4207c-firebase-adminsdk-dtl6f-6878a5d23b.json")
	var config = &firebase.Config{ProjectID: "smart-scale-4207c"}

	// Setup the FireStore data
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, ctx, err
	}

	return client, ctx, nil
}
