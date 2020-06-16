package main

import (
	"github.com/SummerCedrus/ServerKit/netkit"
)

const (
	address = ":50051"
)

func main() {
	netkit.NewRpcServer(address)
}
