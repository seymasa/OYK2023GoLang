package main_test

import (
	m "cashregister"
	"testing"
)

func TestTotalPrice(t *testing.T) {
	want := 15.00
	item := []m.Describable{
		m.Item{"Cips", 15.00, 0.00},
	}
	got := m.TotalPrice(item)

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestTotalPriceEmpty(t *testing.T) {
	want := 0.0
	got := m.TotalPrice([]m.Describable{})
	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
