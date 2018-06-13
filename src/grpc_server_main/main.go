package main

import (
	"netkit"
)

const (
	address = ":50051"
)

func main() {
	netkit.NewRpcServer(address)
}
