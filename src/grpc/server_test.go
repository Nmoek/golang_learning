package grpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	gs := grpc.NewServer()
	us := &Server{}
	RegisterUserServiceServer(gs, us)

	l, err := net.Listen("tcp", ":8090")
	assert.NoError(t, err)

	t.Log("listening....")
	err = gs.Serve(l)
	t.Log(err)
}

func TestClient(t *testing.T) {

	cc, err := grpc.Dial("localhost:8090",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	client := NewUserServiceClient(cc)

	resp, err := client.GetById(context.Background(), &GetByIdRequest{
		Id: 123,
	})
	require.NoError(t, err)

	t.Log("获得远端调用执行的结果", resp.User)

}

// @func: TestOneOf
// @date: 2024-01-04 21:50:54
// @brief: oneof编译后产物用法
// @author: Kewin Li
// @param t
func TestOneOf(t *testing.T) {
	u := &User{
		Lianxi: &User_Email{
			Email: "123@qq.com",
		},
	}

	e, ok := u.Lianxi.(*User_Email)
	if ok {
		t.Log("email", e.Email)

	}

	u = &User{
		Lianxi: &User_Phone{
			Phone: "123456",
		},
	}

	p, ok := u.Lianxi.(*User_Phone)
	if ok {
		t.Log("phone", p.Phone)

	}
}
