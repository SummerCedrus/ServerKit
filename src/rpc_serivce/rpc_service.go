//service function defined here for the server create by netkit.NewRpcServer
package rpc_serivce

import (
	"golang.org/x/net/context"
	."protocol"
	"errors"
)
//server is used to implement rpcSeviceClient
type Server struct{}

// to implement rpcSeviceClient
func (s *Server) Add(ctx context.Context, in *CalParam) (*CalResult, error) {
	result := in.A + in.B
	return &CalResult{Result:result}, nil
}
func (s *Server) Sub(ctx context.Context, in *CalParam) (*CalResult, error) {
	result := in.A - in.B
	return &CalResult{Result:result}, nil
}
func (s *Server) Mul(ctx context.Context, in *CalParam) (*CalResult, error) {
	result := in.A * in.B
	return &CalResult{Result:result}, nil
}
func (s *Server) Div(ctx context.Context, in *CalParam) (*CalResult, error) {
	if 0 == in.B {
		err := errors.New("The divisor cannot be zero!")
		return nil, err
	}
	result := in.A / in.B
	return &CalResult{Result:result}, nil
}