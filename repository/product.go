package repository

import (
	"database/sql"
	"sikm-marketplace-db/model"
)

type ProductRepo interface {
	FindAll() ([]model.Product, error)
	FindByID(id int) (*model.Product, error)
	Save(product *model.Product) error
	Update(id int, product *model.Product) error
}

type productRepoImpl struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepoImpl {
	return &productRepoImpl{
		db: db,
	}
}

func (r *productRepoImpl) FindAll() ([]model.Product, error) {
	return []model.Product{}, nil
}

func (r *productRepoImpl) FindByID(id int) (*model.Product, error) {
	return nil, nil
}

func (r *productRepoImpl) Save(product *model.Product) error {
	return nil
}

func (r *productRepoImpl) Update(id int, product *model.Product) error {
	return nil
}
