package main

import (
	"log"
	"fmt"
	"nutriasoft.com/coypu/msg"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	"context"
	"time"
)


const (
	address     = "localhost:8089"
	defaultPair = "BTC-GBP"
)


func main() {
	fmt.Println("Sending request...")

	test := &msg.CoypuRequest {
		Type: msg.CoypuRequest_BOOK_SNAPSHOT_REQUEST,
		Message: &msg.CoypuRequest_Snap {
			Snap: &msg.BookSnapshot {
				Key: defaultPair,
				Source: 1,
				Levels: 5,
			},
		},
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := msg.NewCoypuServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	r, err := c.RequestData(ctx, test)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Type == msg.CoypuMessage_BOOK_SNAP {
		snap := r.GetSnap()

		bidBook := snap.GetBid()
		for _, level := range bidBook {
			log.Printf(" Bid %f @ %f", level.Qty, level.Px)
		}

		askBook := snap.GetAsk()
		for _, level := range askBook {
			log.Printf(" Ask %f @ %f", level.Qty, level.Px)
		}

	} else {
		log.Printf("Coypu Request Type: %d", r.Type)
		fmt.Println(proto.MarshalTextString(r))
	}
}
