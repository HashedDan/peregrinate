package main

import (
	"log"
	"net"

	pb "github.com/hasheddan/peregrinate/attraction-service/proto/attraction"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func (s *service) CreateAttraction(ctx context.Context, req *pb.Attraction) (*pb.Response, error) {

	attraction, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Attraction: attraction}, nil
}

func (s *service) GetAttractions(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	attractions := s.repo.GetAll()
	return &pb.Response{Attractions: attractions}, nil
}

func main() {
	repo := &Repository{}

	// Set-up gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Connect implementation to grpc service
	pb.RegisterAttractionServiceServer(s, &service{repo})

	// Register reflection service on gRPC server
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
