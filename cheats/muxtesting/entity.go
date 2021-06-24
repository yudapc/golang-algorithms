package muxtesting

import "github.com/Pallinder/go-randomdata"

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price  int64  `json:"price"`
	Status bool   `json:"status"`
}

type Products []Product

func GetProducts() *Products {
	products := &Products{
		{
			ID:     randomdata.Number(10, 20),
			Name:   randomdata.FullName(randomdata.Female),
			Price:  int64(randomdata.Number(1000, 100000)),
			Status: true,
		},
		{
			ID:     randomdata.Number(10, 20),
			Name:   randomdata.FullName(randomdata.Female),
			Price:  int64(randomdata.Number(1000, 100000)),
			Status: true,
		},
		{
			ID:     randomdata.Number(10, 20),
			Name:   randomdata.FullName(randomdata.Female),
			Price:  int64(randomdata.Number(1000, 100000)),
			Status: true,
		},
		{
			ID:     randomdata.Number(10, 20),
			Name:   randomdata.FullName(randomdata.Female),
			Price:  int64(randomdata.Number(1000, 100000)),
			Status: true,
		},
		{
			ID:     randomdata.Number(10, 20),
			Name:   randomdata.FullName(randomdata.Female),
			Price:  int64(randomdata.Number(1000, 100000)),
			Status: true,
		},
	}

	return products
}
