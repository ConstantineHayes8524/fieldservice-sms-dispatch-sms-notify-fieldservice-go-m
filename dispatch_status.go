package main

import "strconv"

type Dispatch struct {
	WorkOrderID string
	PhotoCount  int
	Technician  string
	Status      string
}

func AlertText(d Dispatch) string {
	return "Work order " + d.WorkOrderID + " is " + d.Status + "; " + d.Technician + " has " + strconv.Itoa(d.PhotoCount) + " photo(s)"
}
