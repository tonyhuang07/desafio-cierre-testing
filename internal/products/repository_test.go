package products

import (
	"errors"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller(t *testing.T) {
	t.Run("should return all products by seller", func(t *testing.T) {
		// Given
		products := []Product{
			{
				ID:          "1",
				SellerID:    "1",
				Description: "Product 1",
				Price:       100,
			},
		}
		mock := &store.Mock{
			Data:  []byte(`[{"ID":"1","SellerID":"1","Description":"Product 1","Price":100}, {"ID":"2","SellerID":"2","Description":"Product 2","Price":200}, {"ID":"3","SellerID":"3","Description":"Product 3","Price":300}]`),
			Error: nil,
		}
		db := store.NewFileStore(store.FileType, "products.json")
		db.(*store.FileStore).Mock = mock
		repo := NewRepository(db)
		// When
		result, err := repo.GetAllBySeller("1")
		// Then
		assert.Nil(t, err)
		assert.Equal(t, products, result)
	})
	t.Run("should return nil if id doesn't exist", func(t *testing.T) {
		// Given
		mock := &store.Mock{
			Data:  []byte(`[{"ID":"1","SellerID":"1","Description":"Product 1","Price":100}, {"ID":"2","SellerID":"2","Description":"Product 2","Price":200}, {"ID":"3","SellerID":"3","Description":"Product 3","Price":300}]`),
			Error: nil,
		}
		db := store.NewFileStore(store.FileType, "products.json")
		db.(*store.FileStore).Mock = mock
		repo := NewRepository(db)
		// When
		result, err := repo.GetAllBySeller("4")
		// Then
		assert.Nil(t, err)
		assert.Nil(t, result)
	})
	t.Run("should return error when read file", func(t *testing.T) {
		// Given
		mock := &store.Mock{
			Data:  []byte(`[{ID":"1","SellerID":"1","Description":"Product 1","Price":100}, {"ID":"2","SellerID":"2","Description":"Product 2","Price":200}, {"ID":"3","SellerID":"3","Description":"Product 3","Price":300}]`),
			Error: errors.New("error"),
		}
		db := store.NewFileStore(store.FileType, "products.json")
		db.(*store.FileStore).Mock = mock
		repo := NewRepository(db)
		// When
		result, err := repo.GetAllBySeller("1")
		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
		assert.Nil(t, result)
	})
}
