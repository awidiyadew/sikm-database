package model

import "gorm.io/gorm"

// mapped to table products
type Product struct {
	gorm.Model `json:"-"`
	ID         uint   `gorm:"primarykey;"`
	Name       string // name
	Price      int    // price
	CategoryID int    // category_id
}

// table categories
type Category struct {
	gorm.Model `json:"-"`
	ID         uint   `gorm:"primarykey"`
	Name       string // name
}

type ProductCategory struct {
	ID           int
	Name         string
	Price        int
	CategoryID   int
	CategoryName string
}
