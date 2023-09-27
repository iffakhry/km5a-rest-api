package calculate

import "testing"

func TestAddition(t *testing.T) {
	bil1 := 20
	bil2 := 50
	expected := 70
	actual := Addition(bil1, bil2)
	if actual != expected {
		t.Error("actual & expected tidak sesuai.")
	}
}
