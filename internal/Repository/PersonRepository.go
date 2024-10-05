package repository

import (
	"errors"
	models "go-grpc/internal/Models"
)

type PersonRepository struct {
	id    int
	items map[int]*models.Person
}

func NewRepository() *PersonRepository {
	return &PersonRepository{
		id:    1,
		items: make(map[int]*models.Person, 0),
	}
}

func (obj *PersonRepository) Create(item *models.Person) (*models.Person, error) {
	item.Id = obj.id
	obj.items[obj.id] = item
	obj.id++
	return item, nil
}

func (obj *PersonRepository) GetById(id int) (*models.Person, error) {
	item, ok := obj.items[id]
	if !ok {
		return nil, errors.New("could not found the item")
	}
	return item, nil
}

func (obj *PersonRepository) Delete(id int) error {
	delete(obj.items, id)
	return nil
}

func (obj *PersonRepository) Update(item *models.Person) error {
	obj.items[item.Id] = item
	return nil
}

func (obj *PersonRepository) Search(query string) ([]*models.Person, error) {
	var items []*models.Person
	for _, item := range obj.items {
		if item.Name == query || item.Address == query {
			items = append(items, item)
		}
	}
	return items, nil
}
