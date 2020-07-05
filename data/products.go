package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// Product holds the product data
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` //omit from output
	UpdatedOn   string  `json:"-"` //omit from output
	DeletedOn   string  `json:"-"` //omit from output
}

// GetProducts returns arrays of products
func GetProducts() Products {
	return productList
}

// AddProduct adds the values to the required params
func AddProduct(p *Product)  {
	p.ID=getNextID()
	productList=append(productList,p)
}

// UpdateProduct updates the value of ID
func UpdateProduct(id int, p *Product) error {
	_,pos,err:=findProduct(id)
	if err != nil {
		return err
	}
	p.ID=id
	productList[pos]=p
	return nil
}

func findProduct(id int) (*Product,int,error)  {
	for i, p := range productList {
		if p.ID == id{
			return p,i,nil
		}
	}
	return nil,0,errors.New("No product found")
	
}

func getNextID() int  {
	lp:=productList[len(productList)-1]
	return lp.ID+1
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
