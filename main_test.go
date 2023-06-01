package main

import "testing"

func TestSimpleFactory(t *testing.T) {
    f := Simple("http://localhost:4444")

    if f.Url != "http://localhost:4444" {
        t.Errorf("feature incorrect, got %s, want: %s ", f.Url, "http://localhost:4444")
    }
}