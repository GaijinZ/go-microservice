package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Descripion string  `json:"description"`
	Price      float32 `json:"price"`
	SKU        string  `json:"sku"`
	CreatedOn  string  `json:"-"`
	Updatedon  string  `json:"-"`
	Deletedon  string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList

}

var productList = []*Product{
	{
		ID:         1,
		Name:       "Latte",
		Descripion: "Frothy milky coffee",
		Price:      2.45,
		SKU:        "abc323",
		CreatedOn:  time.Now().UTC().String(),
		Updatedon:  time.Now().UTC().String(),
	},
	{
		ID:         2,
		Name:       "Espresso",
		Descripion: "SHort and strong",
		Price:      1.99,
		SKU:        "fdj34",
		CreatedOn:  time.Now().UTC().String(),
		Updatedon:  time.Now().UTC().String(),
	},
}
