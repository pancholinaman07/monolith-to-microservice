package application

type productReadModel interface {
	AllProducts() ([]products.Product, error)
}

type ProductService struct {
	repo      products.Repository
	readModel productReadModel
}

func NewProductsService() ProductService {

}

func (s ProductService) AllProducts() {

}

func (s ProductService) AddProduct() error {

}
