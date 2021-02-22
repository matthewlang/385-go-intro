package main

import (
	"go-intro/rpc/mailbox/mailbox"

	"flag"
	"fmt"
	"net/rpc"
)

func send(to int, msg string, client *rpc.Client) {
	req := &mailbox.PutMessageRequest{
		UserId:  to,
		Message: msg,
	}
	resp := &mailbox.PutMessageResponse{}

	err := client.Call("MailboxService.PutMessage", req, resp)
	if err != nil {
		fmt.Printf("Error sending: %v\n", err)
		return
	}

	fmt.Println("ok")
}

func get(who int, client *rpc.Client) {
	req := &mailbox.GetMessageRequest{UserId: who}
	resp := &mailbox.GetMessageResponse{}

	err := client.Call("MailboxService.GetMessage", req, resp)
	if err != nil {
		fmt.Printf("Error retrieving: %v\n", err)
		return
	}

	fmt.Printf("%v: %v\n", resp.UserName, resp.Message)
}

func add(who int, name string, client *rpc.Client) {
	req := &mailbox.AddMailboxRequest{
		UserId:   who,
		UserName: name,
	}
	resp := &mailbox.AddMailboxResponse{}

	err := client.Call("MailboxService.AddMailbox", req, resp)
	if err != nil {
		fmt.Printf("Error adding user: %v\n", err)
	}

	fmt.Println("ok")
}

func main() {
	var msg string
	var id int
	var op string
	var name string
	var srv string

	flag.IntVar(&id, "id", 0, "id of user to send to/retrieve for")
	flag.StringVar(&msg, "m", "", "message to send")
	flag.StringVar(&op, "op", "g", "g(et), s(end), or (a)dd")
	flag.StringVar(&name, "name", "", "name to add")
	flag.StringVar(&srv, "srv", "127.0.0.1", "server to dial")
	flag.Parse()

	client, err := rpc.DialHTTP("tcp", srv+":8000")
	if err != nil {
		fmt.Printf("Error connecting to host %v: %v\n", srv, err)
	}

	switch op {
	case "g":
		get(id, client)
	case "s":
		send(id, msg, client)
	case "a":
		add(id, name, client)
	default:
		fmt.Println("-op must be g, s, or a")
	}
}
