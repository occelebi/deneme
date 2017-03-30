package main

import "testing"

func TestFoo(t *testing.T) {
	result := Foo()
	if result != "for" {
		t.Errorf("expecting for, got %s", result)
	}
}
func TestLol(t *testing.T) {
	result := Lol()
	if result != "lof" {
		t.Errorf("expecting lor, got %s", result)
	}
}
func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {
		t.Errorf("expecting bar, got %s", result)
	}
}
