package model

type Product struct {
	ID    int
	Name  string
	Price int
}

type Category struct {
	ID   int
	Name string
}

type ProductCategory struct {
	ID       int
	Name     string
	Price    int
	Category Category
}
