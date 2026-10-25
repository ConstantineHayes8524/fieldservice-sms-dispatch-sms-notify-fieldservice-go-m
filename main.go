package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: fieldservice-sms <phone>")
		os.Exit(2)
	}
	client, err := NewClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	d := Dispatch{WorkOrderID: "WO-42", PhotoCount: 3, Technician: "Mina", Status: "follow-up required"}
	result, err := client.SendSMS(os.Args[1], AlertText(d), d.WorkOrderID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("sent message", result.MessageID)
}
