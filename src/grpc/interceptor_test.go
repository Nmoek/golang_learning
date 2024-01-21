package grpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"net"
	"testing"
)

// InterceptorSuite
// @Description: gRPC拦截器处理
type InterceptorSuite struct {
	suite.Suite
}

func (i *InterceptorSuite) TestServer() {
	t := i.T()
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(NewLogInterceptor(t)))

	RegisterUserServiceServer(server, &Server{
		Name: "interceptor_test",
	})

	l, err := net.Listen("tcp", "localhost:8090")
	assert.NoError(t, err)

	err = server.Serve(l)
	assert.NoError(t, err)

}

func NewLogInterceptor(t *testing.T) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		t.Log("请求处理前:", req, info)
		resp, err = handler(ctx, req)
		t.Log("请求处理后:", resp, err)
		return
	}
}

func TestInterceptor(t *testing.T) {
	suite.Run(t, &InterceptorSuite{})
}
