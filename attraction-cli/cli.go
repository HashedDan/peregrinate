package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/hasheddan/peregrinate/attraction-service/proto/attraction"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "attraction.json"
)

func parseFile(file string) (*pb.Attraction, error) {
	var attraction *pb.Attraction
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &attraction)
	return attraction, err
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAttractionServiceClient(conn)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateAttraction(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetAttractions(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list attractions: %v", err)
	}
	for _, v := range getAll.Attractions {
		log.Println(v)
	}
}
