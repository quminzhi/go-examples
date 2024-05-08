package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/quminzhi/go-examples/grpc-example/route"
	"google.golang.org/grpc"
)

func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln("Failed to get feature:", err)
	}
	fmt.Println(feature)
}

func runSecond(client pb.RouteGuideClient) {
	serverStream, err := client.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{
			Longitude: 121358540,
			Latitude:  313374060,
		},
		Hi: &pb.Point{
			Longitude: 121598790,
			Latitude:  311034130,
		},
	},
	)
	if err != nil {
		log.Fatalln("Failed to list features:", err)
	}

	for {
		feature, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Failed to receive features:", err)
		}
		fmt.Println(feature)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Failed to connect to server:", err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	// runFirst(client)
	runSecond(client)
}
