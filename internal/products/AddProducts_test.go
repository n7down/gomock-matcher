package products

import (
	"fmt"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func EqProduct(p Product) gomock.Matcher {
	return eqProductMatcher{
		product: p,
	}
}

type eqProductMatcher struct {
	product Product
}

func (e eqProductMatcher) Matches(x interface{}) bool {
	p, ok := x.(Product)
	if !ok {
		return false
	}

	if e.product.ID != p.ID {
		return false
	}

	if e.product.Description != p.Description {
		return false
	}

	if e.product.Quantity != p.Quantity {
		return false
	}

	return true
}

func (e eqProductMatcher) String() string {
	return fmt.Sprintf("id: %s description: %s quantity: %d", e.product.ID, e.product.Description, e.product.Quantity)
}

func TestAddProduct_InsertsProduct_ReturnNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := NewMockRepository(mockCtrl)

	expectedProduct := Product{
		ID:          "1",
		Description: "description",
		Quantity:    1,
		CreatedAt:   time.Now(),
	}

	mockRepository.EXPECT().
		Insert(EqProduct(expectedProduct)).
		Return(nil)

	productRequest := AddProductRequest{
		ID:          "1",
		Description: "description",
		Quantity:    1,
	}

	s := NewProductService(mockRepository)
	err := s.AddProduct(productRequest)

	assert.Nil(t, err)
}
