package internal_services

import (
	"context"
	"encoding/json"
	models "go-grpc/internal/Models"
	repository "go-grpc/internal/Repository"
	internal_grpc "go-grpc/internal/grpc"
)

type PersonService struct {
	internal_grpc.UnimplementedPersonServiceServer
	Repo *repository.PersonRepository
}

func (obj *PersonService) Create(context context.Context, in *internal_grpc.CreateRequest) (*internal_grpc.CreateResponse, error) {

	person := &models.Person{
		Name:    in.Name,
		Address: in.Address,
		Salary:  in.Salary,
	}
	model, err := obj.Repo.Create(person)
	if err != nil {
		return &internal_grpc.CreateResponse{
			Message: err.Error(),
		}, nil
	}
	msg, err := json.Marshal(model)
	if err != nil {
		return &internal_grpc.CreateResponse{
			Message: err.Error(),
		}, nil
	}
	return &internal_grpc.CreateResponse{
		Message: string(msg),
	}, nil
}
func (obj *PersonService) GetByID(context context.Context, in *internal_grpc.GetByIdRequest) (*internal_grpc.GetByIdResponse, error) {

	item, err := obj.Repo.GetById(int(in.Id))
	if err != nil {
		return &internal_grpc.GetByIdResponse{
			Result: nil,
		}, err
	}
	return &internal_grpc.GetByIdResponse{
		Result: &internal_grpc.Person{
			Id:      int32(item.Id),
			Name:    item.Name,
			Address: item.Address,
			Salary:  item.Salary,
		},
	}, nil
}
func (obj *PersonService) Update(context context.Context, in *internal_grpc.UpdateRequest) (*internal_grpc.UpdateResponse, error) {
	item, err := obj.Repo.GetById(int(in.Id))
	if err != nil {
		return nil, err
	}
	err = obj.Repo.Update(item)
	if err != nil {
		return nil, err
	}
	return &internal_grpc.UpdateResponse{
		Message: "success",
	}, nil
}
func (obj *PersonService) Delete(context context.Context, in *internal_grpc.DeleteRequest) (*internal_grpc.DeleteResponse, error) {

	err := obj.Repo.Delete(int(in.Id))
	if err != nil {
		return nil, err
	}
	return &internal_grpc.DeleteResponse{
		Message: "success",
	}, nil
}
func (obj *PersonService) Searach(context context.Context, in *internal_grpc.SearachRequest) (*internal_grpc.SearachResponse, error) {
	result, err := obj.Repo.Search(in.Query)
	if err != nil {
		return nil, err
	}
	var items []*internal_grpc.Person
	for _, res := range result {
		items = append(items, &internal_grpc.Person{
			Id:      int32(res.Id),
			Name:    res.Name,
			Address: res.Address,
			Salary:  res.Salary,
		})
	}
	return &internal_grpc.SearachResponse{
		Result: items,
	}, nil
}
