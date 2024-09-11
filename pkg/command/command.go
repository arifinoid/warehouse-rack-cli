package command

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/arifinoid/warehouse-rack-cli/pkg/rack"
)

func ProcessCommand(r *rack.RackSystem, cmd string) *rack.RackSystem {
	args := strings.Fields(cmd)
	if len(args) == 0 {
		return r
	}

	switch args[0] {
	case "create_warehouse_rack":
		if len(args) != 2 {
			fmt.Println("Invalid input for create_warehouse_rack. Usage: create_warehouse_rack <n>")
			return r
		}
		numSlots, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invalid slot number: %s\n", args[1])
			return r
		}
		r = rack.NewRackSystem(numSlots)
		return r
	case "rack":
		if len(args) != 3 {
			fmt.Println("Invalid input for rack. Usage: rack <sku> <exp_date>")
			return r
		}
		r.RackIn(args[1], args[2])
	case "rack_out":
		if len(args) != 2 {
			fmt.Println("Invalid input for rack_out. Usage: rack_out <slot>")
			return r
		}
		slot, err := strconv.Atoi(args[1])
		if err != nil || slot < 0 || slot >= len(r.Slots) {
			fmt.Printf("Invalid slot number: %s\n", args[1])
			return r
		}
		r.RackOut(slot - 1)
	case "status":
		r.Status()
	case "sku_numbers_for_product_with_exp_date":
		if len(args) != 2 {
			fmt.Println("Invalid input for sku_numbers_for_product_with_exp_date. Usage: sku_numbers_for_product_with_exp_date <exp_date>")
			return r
		}
		r.FindSKUByExpDate(args[1])
	case "slot_numbers_for_product_with_exp_date":
		if len(args) != 2 {
			fmt.Println("Invalid input for slot_numbers_for_product_with_exp_date. Usage: slot_numbers_for_product_with_exp_date <exp_date>")
			return r
		}
		r.FindSlotsByExpDate(args[1])
	case "slot_number_for_sku_number":
		if len(args) != 2 {
			fmt.Println("Invalid input for slot_number_for_sku_number. Usage: slot_number_for_sku_number <sku_number>")
			return r
		}
		r.FindSlotBySKU(args[1])
	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Printf("unknown command: %s\n", args[0])
	}

	return r
}
