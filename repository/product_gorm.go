package repository

import (
	"errors"
	"sikm-marketplace-db/model"

	"gorm.io/gorm"
)

type productGormRepoImpl struct {
	db *gorm.DB
}

func NewProductGormRepo(db *gorm.DB) *productGormRepoImpl {
	return &productGormRepoImpl{
		db: db,
	}
}

func (r *productGormRepoImpl) FindAll() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return []model.Product{}, nil
	}
	return products, nil
}

func (r *productGormRepoImpl) FindByID(id int) (*model.Product, error) {
	var prd model.Product
	result := r.db.Where("id = ?", id).First(&prd)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		return nil, result.Error
	}
	return &prd, nil
}

func (r *productGormRepoImpl) Save(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *productGormRepoImpl) Update(id int, product *model.Product) error {
	result := r.db.Where("id = ?", id).Updates(product)
	if result.RowsAffected == 0 {
		return errors.New("no product updated")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *productGormRepoImpl) FindDetail(id int) (*model.Product, error) {
	var prd model.Product
	result := r.db.Preload("Category").First(&prd, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &prd, nil
}
