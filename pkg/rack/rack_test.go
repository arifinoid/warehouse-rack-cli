package rack

import (
	"testing"

	"github.com/arifinoid/warehouse-rack-cli/lib"
)

func TestRackSystem(t *testing.T) {
	r := NewRackSystem(6)

	expectedSlots := 6
	if len(r.Slots) != expectedSlots {
		t.Errorf("Expected %d slots, got %d", expectedSlots, len(r.Slots))
	}

	// rack-in
	r.RackIn("ZG11AQA", "2024-02-28")
	if r.Slots[0].SKU != "ZG11AQA" || r.Slots[0].ExpiryDate != "2024-02-28" {
		t.Errorf("Expected SKU ZG11AQA in slot 1, got %v", r.Slots[0].SKU)
	}

	r.RackIn("SD92349WW", "2024-02-28")
	if r.Slots[1].SKU != "SD92349WW" || r.Slots[1].ExpiryDate != "2024-02-28" {
		t.Errorf("Expected SKU SD92349WW in slot 2, got %v", r.Slots[1].SKU)
	}

	// rack-out
	r.RackOut(0)
	if r.Slots[0] != nil {
		t.Errorf("Expected slot 1 to be empty after checkout")
	}

	// rack-out: test an already free slot (expecting a message)
	output := lib.CaptureOutput(func() {
		r.RackOut(0)
	})

	expectedMessage := "Invalid slot or slot is already free\n"
	if output != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, output)
	}

	// rack_out: test invalid on a non-existent slot (slot index 3)
	output = lib.CaptureOutput(func() {
		r.RackOut(3)
	})
	if output != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, output)
	}

	// status
	output = lib.CaptureOutput(func() {
		r.Status()
	})
	if output != "Slot No.\tSKU No.\t\tExp Date\n2\t\tSD92349WW\t\t2024-02-28\n" {
		t.Errorf("Expected message %q, got %q", "Slot No.\tSKU No.\t\tExp Date\n2\t\tSD92349WW\t\t2024-02-28\n", output)
	}

	// sku_numbers_for_product_with_exp_date
	output = lib.CaptureOutput(func() {
		r.FindSKUByExpDate("2024-02-28")
	})
	if output != "SD92349WW\n" {
		t.Errorf("Expected message %q, got %q", "SD92349WW\n", output)
	}

	// slot_numbers_for_product_with_exp_date
	output = lib.CaptureOutput(func() {
		r.FindSlotsByExpDate("2024-02-28")
	})
	if output != "2\n" {
		t.Errorf("Expected message %q, got %q", "2\n", output)
	}

	// slot_number_for_sku_number
	output = lib.CaptureOutput(func() {
		r.FindSlotBySKU("SD92349WW")
	})
	if output != "2\n" {
		t.Errorf("Expected message %q, got %q", "2\n", output)
	}
}
