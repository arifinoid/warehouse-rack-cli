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

func (r *RackSystem) FindNearestSlot() int {
	for i := 0; i < len(r.Slots); i++ {
		if r.Slots[i] == nil {
			return i
		}
	}
	return -1
}

func (r *RackSystem) RackIn(sku string, expDate string) {
	slot := r.FindNearestSlot()
	if slot == -1 {
		fmt.Println("Sorry, rack is full")
		return
	}
	r.Slots[slot] = &Product{SKU: sku, ExpiryDate: expDate}
	fmt.Printf("Allocated slot number: %d\n", slot+1)
}

func (r *RackSystem) RackOut(slot int) {
	if slot < 0 || slot >= len(r.Slots) || r.Slots[slot] == nil {
		fmt.Println("Invalid slot or slot is already free")
		return
	}
	r.Slots[slot] = nil
	fmt.Printf("Slot number %d is free\n", slot+1)
}
