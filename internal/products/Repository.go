//go:generate mockgen -destination MockResository.go -source Repository.go -package products
package products

type Repository interface {
	Insert(Product) error
}
