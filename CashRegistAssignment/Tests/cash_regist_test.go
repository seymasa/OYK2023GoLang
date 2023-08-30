package main

import (
	"testing"
)

func TestBıdıbıdı(t *testing.T) {
	want := "hi vigo!"
	got := ""

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
