package internal_services

import (
	"context"
	internal_grpc "go-grpc/internal/grpc"
)

type Person struct {
	Id      int
	Name    string
	Address string
	Salary  float32
}

type Repository struct {
	id    int
	items []*Person
}

func NewRepository() *Repository {
	return &Repository{
		id:    1,
		items: make([]*Person, 0),
	}
}

type PersonService struct {
	internal_grpc.UnimplementedPersonServiceServer
	repo *Repository
}

func (obj *PersonService) Create(context context.Context, in *internal_grpc.CreateRequest) (*internal_grpc.CreateResponse, error) {

	persons = append(persons, nil)
	return &internal_grpc.CreateResponse{}, nil
}
func (obj *PersonService) GetByID(context context.Context, _ *internal_grpc.GetByIdRequest) (*internal_grpc.GetByIdResponse, error) {

	return &internal_grpc.GetByIdResponse{}, nil
}
func (obj *PersonService) Update(context context.Context, _ *internal_grpc.UpdateRequest) (*internal_grpc.UpdateResponse, error) {
	return &internal_grpc.UpdateResponse{}, nil
}
func (obj *PersonService) Delete(context context.Context, _ *internal_grpc.DeleteRequest) (*internal_grpc.DeleteResponse, error) {
	return &internal_grpc.DeleteResponse{}, nil
}
func (obj *PersonService) Searach(context context.Context, _ *internal_grpc.SearachRequest) (*internal_grpc.SearachResponse, error) {
	return &internal_grpc.SearachResponse{}, nil
}
