package rack

import (
	"testing"
)

func TestRackSystem(t *testing.T) {
	r := NewRackSystem(6)

	expectedSlots := 6
	if len(r.Slots) != expectedSlots {
		t.Errorf("Expected %d slots, got %d", expectedSlots, len(r.Slots))
	}
}
