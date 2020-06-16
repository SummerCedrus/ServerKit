package main

import (
	"github/SummerCedrus/ServerKit/netkit"
)

const (
	address = ":50051"
)

func main() {
	netkit.NewRpcServer(address)
}
