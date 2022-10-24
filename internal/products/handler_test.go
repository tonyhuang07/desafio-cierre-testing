package products

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var s = createService()

func createService() *gin.Engine {
	db := store.NewFileStore(store.FileType, "products.json")
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	r := gin.Default()
	r.GET("/products", handler.GetProducts)
	return r
}

func createWrongService() *gin.Engine {
	db := store.NewFileStore(store.FileType, "")
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	r := gin.Default()
	r.GET("/products", handler.GetProducts)
	return r
}

func createRequestTest(method, path string) *http.Request {
	req, _ := http.NewRequest(method, path, nil)
	return req
}

func TestGetProducts(t *testing.T) {
	t.Run("should return all products", func(t *testing.T) {
		path := fmt.Sprintf("/products?seller_id=%s", "1")
		req := createRequestTest("GET", path)
		w := httptest.NewRecorder()
		s.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 400 if query param is null", func(t *testing.T) {
		path := fmt.Sprintf("/products?seller_id=%s", "")
		req := createRequestTest("GET", path)
		w := httptest.NewRecorder()
		s.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return 500 if file doesn't exist", func(t *testing.T) {
		wrongService := createWrongService()
		path := fmt.Sprintf("/products?seller_id=%s", "2")
		req := createRequestTest("GET", path)
		w := httptest.NewRecorder()
		wrongService.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
