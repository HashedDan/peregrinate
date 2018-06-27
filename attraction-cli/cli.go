package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/hasheddan/peregrinate/attraction-service/proto/attraction"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

const (
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

	cmd.Init()

	// Create new greeter client
	client := pb.NewAttractionServiceClient("go.micro.srv.attraction", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	attraction, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateAttraction(context.TODO(), attraction)
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
