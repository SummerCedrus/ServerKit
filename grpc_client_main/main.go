
package main

import (
	"log"
	."github.com/SummerCedrus/ServerKit/protocol"
	"github.com/SummerCedrus/ServerKit/netkit"
	"time"
	"golang.org/x/net/context"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	c := netkit.NewRpcClient(address)
	if nil == c {
		return
	}
	for  {
		time.Sleep(5*time.Second)
		r, err := c.Add(context.Background(), &CalParam{A:3,B:2})
		if err != nil {
			log.Printf("could not cal: %v", err)
			continue
		}
		log.Printf("Result: %d", r.GetResult())
	}

}

