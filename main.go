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
	defaultPair = "BTC/USD"
)


func main() {
	fmt.Println("Hello, world.")

	test := &msg.CoypuRequest {
		Type: msg.CoypuRequest_BOOK_SNAPSHOT_REQUEST,
		Message: &msg.CoypuRequest_Snap {
			Snap: &msg.BookSnapshot {
				Key: defaultPair,
				Source: 1,
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
	log.Printf("Coypu Request Type: %d", r.Type)
	fmt.Println(proto.MarshalTextString(r))
}
