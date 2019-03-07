package main

import (
	"log"
	"fmt"
	"nutriasoft.com/coypu/msg"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	"context"
	"time"
	"os"
)

const (
	//address     = "localhost:8089"
	address = "35.185.73.107:80"
)


func main() {
	fmt.Println("Sending request...")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := msg.NewCoypuServiceClient(conn)

	var i uint32 = 0
	for i = 1; i <= 10; i++ {
		test := &msg.CoypuRequest {
			Type: msg.CoypuRequest_BOOK_SNAPSHOT_REQUEST,
			Message: &msg.CoypuRequest_Snap {
				Snap: &msg.BookSnapshot {
					Key: os.Args[1],
					Source: 1,
					Levels: i,
				},
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		r, err := c.RequestData(ctx, test)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if r.Type == msg.CoypuMessage_BOOK_SNAP {
			snap := r.GetSnap()
			
			bidBook := snap.GetBid()
			for x, level := range bidBook {
				log.Printf(" Bid %2d. %14.8f @ %-14.4f", x+1, level.Qty/100000000.0, level.Px/100000000.0)
			}
			
			askBook := snap.GetAsk()
			for x, level := range askBook {
				log.Printf(" Ask %2d. %14.8f @ %-14.4f", x+1, level.Qty/100000000.0, level.Px/100000000.0)
			}
			
		} else {
			log.Printf("Coypu Request Type: %d", r.Type)
			fmt.Println(proto.MarshalTextString(r))
		}
		time.Sleep(time.Duration(i) * 200 * time.Millisecond)
	}
}
