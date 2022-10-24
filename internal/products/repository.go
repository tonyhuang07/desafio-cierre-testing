package products

import "github.com/bootcamp-go/desafio-cierre-testing/pkg/store"

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return nil, err
	}

	var filteredProducts []Product
	for _, product := range products {
		if product.SellerID == sellerID {
			filteredProducts = append(filteredProducts, product)
		}
	}

	return filteredProducts, nil
}
