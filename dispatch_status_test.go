package main

import "testing"

func TestAlertText(t *testing.T) {
	d := Dispatch{WorkOrderID: "WO-42", PhotoCount: 3, Technician: "Mina", Status: "follow-up required"}
	want := "Work order WO-42 is follow-up required; Mina has 3 photo(s)"
	if got := AlertText(d); got != want {
		t.Fatalf("AlertText() = %q, want %q", got, want)
	}
}
