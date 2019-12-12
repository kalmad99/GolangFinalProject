package menu

import "../entity"

// CategoryRepository specifies menu category related database operations
type ProductRepository interface {
	Products() ([]entity.Product, error)
	MobProducts() ([]entity.Product, error)
	CamProducts() ([]entity.Product, error)
	CompProducts() ([]entity.Product, error)
	LapProducts() ([]entity.Product, error)
	Product(id int) (entity.Product, error)
	SearchProduct(index string) ([]entity.Product, error)
	UpdateProduct(product entity.Product) error
	DeleteProduct(id int) error
	StoreProduct(product entity.Product) error
	RateProduct(product entity.Product) (entity.Product, error)
}

type UserRepository interface {
	Users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}
