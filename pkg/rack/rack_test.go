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
}
