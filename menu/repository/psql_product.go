package repository

import (
	"../../entity"
	"database/sql"
	"errors"
	"math"
)

// PsqlProductRepository implements the
// menu.ProductRepository interface
type PsqlProductRepository struct {
	conn *sql.DB
}

// NewPsqlProductRepository will create an object of PsqlProductRepository
func NewPsqlProductRepository(Conn *sql.DB) *PsqlProductRepository {
	return &PsqlProductRepository{conn: Conn}
}

// Products returns all products from the database
func (pr *PsqlProductRepository) Products() ([]entity.Product, error) {

	rows, err := pr.conn.Query("SELECT * FROM products;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	prds := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.ItemType,
			&product.Quantity, &product.Price, &product.Description, &product.Image,
			&product.Rating, &product.RatersCount)
		if err != nil {
			return nil, err
		}
		prds = append(prds, product)
	}

	return prds, nil
}

func (pr *PsqlProductRepository) MobProducts() ([]entity.Product, error) {

	rows, err := pr.conn.Query("SELECT * FROM products where itemtype=?;", "mobile")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	prds := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.ItemType,
			&product.Quantity, &product.Price, &product.Description, &product.Image,
			&product.Rating, &product.RatersCount)
		if err != nil {
			return nil, err
		}
		prds = append(prds, product)
	}

	return prds, nil
}
func (pr *PsqlProductRepository) CamProducts() ([]entity.Product, error) {

	rows, err := pr.conn.Query("SELECT * FROM products where itemtype=?;", "camera")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	prds := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.ItemType,
			&product.Quantity, &product.Price, &product.Description, &product.Image,
			&product.Rating, &product.RatersCount)
		if err != nil {
			return nil, err
		}
		prds = append(prds, product)
	}

	return prds, nil
}
func (pr *PsqlProductRepository) CompProducts() ([]entity.Product, error) {

	rows, err := pr.conn.Query("SELECT * FROM products where itemtype=?;", "computer")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	prds := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.ItemType,
			&product.Quantity, &product.Price, &product.Description, &product.Image,
			&product.Rating, &product.RatersCount)
		if err != nil {
			return nil, err
		}
		prds = append(prds, product)
	}

	return prds, nil
}
func (pr *PsqlProductRepository) LapProducts() ([]entity.Product, error) {

	rows, err := pr.conn.Query("SELECT * FROM products where itemtype=?;", "laptop")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	prds := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.ItemType,
			&product.Quantity, &product.Price, &product.Description, &product.Image,
			&product.Rating, &product.RatersCount)
		if err != nil {
			return nil, err
		}
		prds = append(prds, product)
	}

	return prds, nil
}

// Product returns a product with a given id
func (pr *PsqlProductRepository) Product(id int) (entity.Product, error) {

	row := pr.conn.QueryRow("SELECT * FROM products WHERE id = ?", id)

	p := entity.Product{}

	err := row.Scan(&p.ID, &p.Name, &p.ItemType,
		&p.Quantity, &p.Price, &p.Description, &p.Image, &p.Rating, &p.RatersCount)
	if err != nil {
		return p, err
	}

	return p, nil
}

//Searches for the product
func (pr *PsqlProductRepository) SearchProduct(index string) ([]entity.Product, error) {
	//query := "SELECT * FROM products WHERE itemname LIKE ?"
	//rows, err := pr.conn.Query(query, "'%'" + index + "%'")
	rows, err := pr.conn.Query("SELECT * FROM products WHERE itemname LIKE ?", "%" + index + "%" )
	if err != nil {
		panic(err.Error())
			//errors.New("Could not query the database")
	}
	defer rows.Close()

	prds := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.ItemType,
			&product.Quantity, &product.Price, &product.Description, &product.Image,
			&product.Rating, &product.RatersCount)
		if err != nil {
			return nil, err
		}
		prds = append(prds, product)
	}
	return prds, nil
}

// UpdateProduct updates a given product with a new data
func (pr *PsqlProductRepository) UpdateProduct(c entity.Product) error {

	_, err := pr.conn.Exec("UPDATE products SET itemname=?,itemtype=?," +
		"quantity=?,price=?,description=?,image=? WHERE id=?",
		c.Name, c.ItemType, c.Quantity, c.Price, c.Description, c.Image, c.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteProduct removes a product from a database by its id
func (pr *PsqlProductRepository) DeleteProduct(id int) error {

	_, err := pr.conn.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreProduct stores new product information to database
func (pr *PsqlProductRepository) StoreProduct(c entity.Product) error {

	_, err := pr.conn.Exec("INSERT INTO products (itemname,itemtype,quantity,price,description,image)" +
	" values(?, ?, ?, ?, ?, ?)",c.Name, c.ItemType, c.Quantity, c.Price, c.Description, c.Image)
	if err != nil {
		//panic(err)
		return errors.New("Insertion has failed")
	}

	return nil
}
func (pr *PsqlProductRepository) RateProduct(c entity.Product) (entity.Product, error) {

	var oldratings float64
	var oldcount float64

	_ = pr.conn.QueryRow("SELECT rating, raterscount FROM products WHERE id = ?", c.ID).Scan(&oldratings, &oldcount)

	var newratings = ((oldratings*oldcount) + c.Rating)/(oldcount+1)


	_, err := pr.conn.Exec("UPDATE products SET rating=?,raterscount=? WHERE id=?",
		float64(math.Round(newratings*2))/2, oldcount+1, c.ID)
	if err != nil {
		return c, errors.New("Updating Rate has failed")
	}

	return c, nil
}
