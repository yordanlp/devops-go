package main

import "testing"

func TestSimpleFactory_ok(t *testing.T) {
	f := SimpleFactory("http://localhost")

	if f.Url != "http://localhost" {
		t.Errorf("feature incorrect, got %s, want: %s ", f.Url, "http://localhost")
	}
}

/*func TestSimpleFactory_fail(t *testing.T) {
	f := SimpleFactory("http://localhost")

	if f.Description != "World" {
		t.Errorf("feature incorrect, got %s, want: %s ", f.Description, "World")
	}
}*/
