package productDomain

import (
	"time"
)

type Product struct {
	ID          int
	Code        string
	Name        string
	Description string
	Origin      int
	Qty         int
	Price       int
	Weight      float64
	Status      bool
	Category_Id int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Category struct {
	ID        int
	Name      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
