package calculate

import "testing"

func TestCheckPrime(t *testing.T) {
	t.Run("test bilangan prima 11 = true", func(t *testing.T) {
		number := 11
		actual := CheckPrime(number)
		expected := true
		if actual != expected {
			t.Error("actual & expected tidak sesuai")
		}
	})

	t.Run("test bilangan 1", func(t *testing.T) {
		number := 1
		actual := CheckPrime(number)
		expected := false
		if actual != expected {
			t.Error("actual & expected tidak sesuai")
		}
	})

	t.Run("test bilangan 10", func(t *testing.T) {
		number := 15
		actual := CheckPrime(number)
		expected := false
		if actual != expected {
			t.Error("actual & expected tidak sesuai")
		}
	})
	t.Run("test bilangan 2", func(t *testing.T) {
		number := 2
		actual := CheckPrime(number)
		expected := true
		if actual != expected {
			t.Error("actual & expected tidak sesuai")
		}
	})
}
