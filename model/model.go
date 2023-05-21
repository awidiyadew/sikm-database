package model

type Product struct {
	ID         int
	Name       string
	Price      int
	CategoryID int
}

type Category struct {
	ID   int
	Name string
}
