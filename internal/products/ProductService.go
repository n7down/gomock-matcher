package products

import "time"

type ProductService struct {
	repository Repository
}

func NewProductService(r Repository) *ProductService {
	return &ProductService{
		repository: r,
	}
}

func (p *ProductService) AddProduct(r AddProductRequest) (err error) {
	product := Product{
		ID:          r.ID,
		Description: r.Description,
		Quantity:    r.Quantity,
		CreatedAt:   time.Now(),
	}
	err = p.repository.Insert(product)
	return
}
