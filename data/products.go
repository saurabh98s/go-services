package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product holds the product data
type Product struct {
	// the id for this user
	// 
	// required: true
	// min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name"  validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"` //omit from output
	UpdatedOn   string  `json:"-"` //omit from output
	DeletedOn   string  `json:"-"` //omit from output
}

// GetProducts returns arrays of products
func GetProducts() Products {
	return productList
}

// AddProduct adds the values to the required params
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct updates the value of ID
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, 0, errors.New("No product found")

}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// Products holds an array of products
type Products []*Product

// FromJSON decodes data to JSON
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// ToJSON Encodes data to JSON
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)

}

// Validate checks if input matches the parameters
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", ValidateSKU)
	return validate.Struct(p)
}

func ValidateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)

	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// productList is a collection of products
var productList = []*Product{

	{
		ID:          1,
		Name:        "Latte",
		Description: "frothy milk coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().Local().String(),
		UpdatedOn:   time.Now().Local().String(),
	},
	{

		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee , no milk",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().Local().String(),
		UpdatedOn:   time.Now().Local().String(),
	},
}
