package main

// ProductDB represents the database operations interface
type ProductDB interface {
	GetProduct(id string) (*Product, error)
	SaveProduct(product *Product) error
	GetProductsByCategory(category string) ([]*Product, error)
}

// MockDB is a mock implementation of ProductDB for demo purposes
type MockDB struct {
	products map[string]*Product
}

// NewMockDB creates a new mock database
func NewMockDB() *MockDB {
	return &MockDB{
		products: make(map[string]*Product),
	}
}

// GetProduct retrieves a product by ID
func (db *MockDB) GetProduct(id string) (*Product, error) {
	if product, ok := db.products[id]; ok {
		return product, nil
	}
	return nil, ErrProductNotFound
}

// SaveProduct saves a product
func (db *MockDB) SaveProduct(product *Product) error {
	db.products[product.ID] = product
	return nil
}

// GetProductsByCategory retrieves products by category
func (db *MockDB) GetProductsByCategory(category string) ([]*Product, error) {
	var products []*Product
	for _, p := range db.products {
		if p.Category == category {
			products = append(products, p)
		}
	}
	return products, nil
}
