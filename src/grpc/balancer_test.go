package grpc

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/weightedroundrobin" // 匿名引包: 触发包的初始化
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"net"
	"testing"
	"time"
)

type BalancerTestSuite struct {
	suite.Suite
	cli *etcdv3.Client
}

func (b *BalancerTestSuite) SetupSuite() {

	cli, err := etcdv3.NewFromURL("localhost:2379")
	assert.NoError(b.T(), err)
	b.cli = cli

}

func (b *BalancerTestSuite) TearDownTest() {

}
func (b *BalancerTestSuite) TestServer() {
	go func() {
		b.startServer("127.0.0.1:8090", 10, &Server{
			Name: "127.0.0.1:8090",
		})
	}()

	go func() {
		b.startServer("127.0.0.1:8091", 20, &Server{
			Name: "127.0.0.1:8091",
		})
	}()
	b.startServer("127.0.0.1:8092", 30, &Server{
		Name: "127.0.0.1:8092",
	})
}

func (b *BalancerTestSuite) startServer(addr string, weight int, svc UserServiceServer) {
	t := b.T()

	// 1. 创建gRPC服务端监听接口
	l, err := net.Listen("tcp", addr)
	assert.NoError(t, err)

	// 2. 注册一个新服务节点
	em, err := endpoints.NewManager(b.cli, "service/user")
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 开启租约机制
	ttl := int64(5)
	leaseResp, err := b.cli.Grant(ctx, ttl)
	assert.NoError(t, err)

	key := "service/user/" + addr
	err = em.AddEndpoint(ctx, key, endpoints.Endpoint{

		Addr: addr, // 定位信息
		Metadata: map[string]any{
			"weight": weight,
		},
	}, etcdv3.WithLease(leaseResp.ID)) //添加租约

	assert.NoError(t, err)

	// 增加续约机制
	kaCtx, kaCancel := context.WithCancel(context.Background())
	go func() {
		_, err2 := b.cli.KeepAlive(kaCtx, leaseResp.ID)
		assert.NoError(t, err2)
		//for kaResp := range ch {
		//	t.Log(kaResp.String())
		//}
	}()

	// 3. 开启gRPC服务端
	server := grpc.NewServer()
	RegisterUserServiceServer(server, svc)
	t.Log(addr, "listening...")
	err = server.Serve(l)
	assert.NoError(t, err)

	// 先关闭续约机制
	kaCancel()

	// 4. 删除注册信息
	err = em.DeleteEndpoint(ctx, key)
	assert.NoError(t, err)

	// 5. 优雅关闭gRPC服务器
	server.GracefulStop()

	// 6. 关闭服务注册
	//err = b.cli.Close()
	//assert.NoError(t, err)

}

// @func: TestClientCustom
// @date: 2024-01-21 01:06:19
// @brief: 负载均衡-自定义加权轮询算法
// @author: Kewin Li
// @receiver b
func (b *BalancerTestSuite) TestClientCustomWRR() {

	Init()

	t := b.T()
	resolverEtcd, err := resolver.NewBuilder(b.cli)
	assert.NoError(t, err)

	// 一定要三个'/' !!!
	cc, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(resolverEtcd),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 选择gRPC内置负载均衡算法
		grpc.WithDefaultServiceConfig(`
{
    "loadBalancingConfig": [
        {
            "custom_weighted_round_robin": {}
        }
    ]
}
`))

	assert.NoError(t, err)

	userClient := NewUserServiceClient(cc)

	for {
		id := rand.Int63n(100)
		resp, err2 := userClient.GetById(context.Background(), &GetByIdRequest{
			Id: id,
		})

		assert.NoError(t, err2)

		fmt.Printf("rpc ret: %v \n", resp.User)
		time.Sleep(time.Second * 2)
	}

}

// @func: TestClientWRR
// @date: 2024-01-20 23:17:12
// @brief: 负载均衡-加权轮询算法
// @author: Kewin Li
// @receiver b
func (b *BalancerTestSuite) TestClientWRR() {
	t := b.T()
	resolverEtcd, err := resolver.NewBuilder(b.cli)
	assert.NoError(t, err)

	// 一定要三个'/' !!!
	cc, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(resolverEtcd),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 选择gRPC内置负载均衡算法
		grpc.WithDefaultServiceConfig(`
{
    "loadBalancingConfig": [
        {
            "weighted_round_robin": {}
        }
    ]
}
`))

	assert.NoError(t, err)

	userClient := NewUserServiceClient(cc)

	for {
		id := rand.Int63n(100)
		resp, err2 := userClient.GetById(context.Background(), &GetByIdRequest{
			Id: id,
		})

		assert.NoError(t, err2)

		fmt.Printf("rpc ret: %v \n", resp.User)
		time.Sleep(time.Second * 2)
	}

}

// @func: TestClient
// @date: 2024-01-20 23:16:54
// @brief: 负载均衡-轮询算法
// @author: Kewin Li
// @receiver b
func (b *BalancerTestSuite) TestClient() {
	t := b.T()
	resolverEtcd, err := resolver.NewBuilder(b.cli)
	assert.NoError(t, err)

	// 一定要三个'/' !!!
	cc, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(resolverEtcd),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 选择gRPC内置负载均衡算法
		grpc.WithDefaultServiceConfig(`
{
    "loadBalancingConfig": [
        {
            "round_robin": {}
        }
    ]
}
`))

	assert.NoError(t, err)

	userClient := NewUserServiceClient(cc)

	for {
		id := rand.Int63n(100)
		resp, err2 := userClient.GetById(context.Background(), &GetByIdRequest{
			Id: id,
		})

		assert.NoError(t, err2)

		fmt.Printf("rpc ret: %v \n", resp.User)
		time.Sleep(time.Second * 2)
	}

}

func TestBalancer(t *testing.T) {
	suite.Run(t, &BalancerTestSuite{})
}
