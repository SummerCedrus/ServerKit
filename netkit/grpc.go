//default option to create grpc client and server
package netkit

import (

	."github.com/SummerCedrus/ServerKit/protocol"

	"google.golang.org/grpc"
	"log"
	"net"
	"google.golang.org/grpc/reflection"
	"github.com/SummerCedrus/ServerKit/rpc_serivce"
)
//缺省创建rpc客户端
func NewRpcClient(address string)(RpcSeviceClient){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil
	}

	c := NewRpcSeviceClient(conn)

	return  c
}

//缺省创建rpc服务器
func NewRpcServer(address string){
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterRpcSeviceServer(s, &rpc_serivce.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}