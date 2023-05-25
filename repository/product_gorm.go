package repository

import (
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

// func (r *productGormRepoImpl) FindByID(id int) (*model.Product, error) {
// 	row := r.db.QueryRow("select id, name, price, category_id from products where id = $1", id)
// 	if row.Err() != nil {
// 		return nil, row.Err()
// 	}
// 	var item model.Product
// 	err := row.Scan(&item.ID, &item.Name, &item.Price, &item.CategoryID)
// 	// check sentinel error not found
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return nil, errors.New("product not found")
// 	}
// 	return &item, row.Err()
// }

// func (r *productGormRepoImpl) Save(product *model.Product) error {
// 	// todo: teman2 bisa pake result untuk cek rowsAffected
// 	_, err := r.db.Exec("INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3)", product.Name, product.Price, product.CategoryID)
// 	return err
// }

// func (r *productGormRepoImpl) Update(id int, product *model.Product) error {
// 	_, err := r.db.Exec("UPDATE products SET name=$1, price=$2, category_id=$3 WHERE id=$4", product.Name, product.Price, product.CategoryID, id)
// 	return err
// }

// func (r *productGormRepoImpl) FindDetail(id int) (*model.ProductCategory, error) {
// 	row := r.db.QueryRow("SELECT p.id, p.name, p.price, p.category_id, c.name FROM products p JOIN categories c ON c.id = p.category_id WHERE p.id = $1", id)
// 	var prd model.ProductCategory
// 	if row.Err() != nil {
// 		fmt.Println("failed to get")
// 	}
// 	if err := row.Scan(&prd.ID, &prd.Name, &prd.Price, &prd.CategoryID, &prd.CategoryName); err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, errors.New("product not found")
// 		}
// 		return nil, err
// 	}
// 	return &prd, nil
// }
