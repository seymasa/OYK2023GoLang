package main_test

import (
	"cashregister"
	"testing"
)

func TestTotalPrice(t *testing.T) {
	want := 15.00
	item := []main.Describable{
		main.Item{"Cips", 15.00, 0.00},
	}
	got := main.TotalPrice(item)

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
