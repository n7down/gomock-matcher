package products

import (
	"time"
)

type Product struct {
	ID          string
	Description string
	Quantity    int
	CreatedAt   time.Time
}
