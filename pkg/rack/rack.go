package rack

import "fmt"

type Product struct {
	SKU        string
	ExpiryDate string
}

type RackSystem struct {
	Slots [](*Product)
}

func NewRackSystem(n int) *RackSystem {
	fmt.Printf("Created a warehouse rack with %d slots\n", n)
	return &RackSystem{
		Slots: make([](*Product), n),
	}
}
