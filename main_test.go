package main

import (
	"strings"
	"testing"
)

func TestGreeter(t *testing.T) {
	expected := "Hello Trevor!"
	name := "Trevor"
	res := greet(&name)
	if strings.Compare(res, expected) != 0 {
		t.Errorf("Expected %s for name %s", expected, res)
	}
}

