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
		return r
	case "rack_out":
		return r
	case "status":
		return r
	case "sku_numbers_for_product_with_exp_date":
		return r
	case "slot_numbers_for_product_with_exp_date":
		return r
	case "slot_number_for_sku_number":
		return r
	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Printf("unknown command: %s\n", args[0])
	}

	return r
}
