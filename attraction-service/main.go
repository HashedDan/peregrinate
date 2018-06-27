package main

import (
	"fmt"

	pb "github.com/hasheddan/peregrinate/attraction-service/proto/attraction"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

// IRepository is an interface for the repository type
type IRepository interface {
	Create(*pb.Attraction) (*pb.Attraction, error)
	GetAll() []*pb.Attraction
}

// Repository is a temporary stand-in for a consistent datastore
type Repository struct {
	attractions []*pb.Attraction
}

// Create adds an attraction to the datastore
func (repo *Repository) Create(attraction *pb.Attraction) (*pb.Attraction, error) {
	updated := append(repo.attractions, attraction)
	repo.attractions = updated
	return attraction, nil
}

func (repo *Repository) GetAll() []*pb.Attraction {
	return repo.attractions
}

type service struct {
	repo IRepository
}

func (s *service) CreateAttraction(ctx context.Context, req *pb.Attraction, res *pb.Response) error {

	// Save our consignment
	attraction, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Attraction = attraction
	return nil
}

func (s *service) GetAttractions(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	attractions := s.repo.GetAll()
	res.Attractions = attractions
	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.attraction"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterAttractionServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
