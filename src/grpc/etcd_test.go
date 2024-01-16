package grpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"net"
	"testing"
	"time"
)

type EtcdTestSuite struct {
	suite.Suite
	cli *etcdv3.Client
}

func (e *EtcdTestSuite) SetupSuite() {
	cli, err := etcdv3.NewFromURL("localhost:2379")
	assert.NoError(e.T(), err)
	e.cli = cli
}

func (e *EtcdTestSuite) TearDownTest() {

}

func (e *EtcdTestSuite) TestServer() {
	t := e.T()

	// 1. 创建gRPC服务端监听接口
	l, err := net.Listen("tcp", "localhost:8090")
	assert.NoError(t, err)

	// 2. 注册一个新服务节点
	em, err := endpoints.NewManager(e.cli, "service/user")
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 开启租约机制
	ttl := int64(5)
	leaseResp, err := e.cli.Grant(ctx, ttl)
	assert.NoError(t, err)

	addr := "localhost:8090"
	key := "service/user/" + addr
	err = em.AddEndpoint(ctx, key, endpoints.Endpoint{

		Addr: addr, // 定位信息
		//Metadata: ,  //一些私有数据
	}, etcdv3.WithLease(leaseResp.ID)) //添加租约

	assert.NoError(t, err)

	// 增加续约机制
	kaCtx, kaCancel := context.WithCancel(context.Background())
	go func() {
		ch, err2 := e.cli.KeepAlive(kaCtx, leaseResp.ID)
		assert.NoError(t, err2)
		for kaResp := range ch {
			t.Log(kaResp.String())
		}
	}()

	// 模拟注册信息变动
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for now := range ticker.C {
			ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err2 := em.Update(ctx2, []*endpoints.UpdateWithOpts{
				{
					Update: endpoints.Update{
						Op:  endpoints.Add,
						Key: key,
						Endpoint: endpoints.Endpoint{
							Addr:     addr,
							Metadata: now.String(),
						},
					},
					// 加上租约
					Opts: []etcdv3.OpOption{
						etcdv3.WithLease(leaseResp.ID),
					},
				},
			})
			// 以上一连串等同于下面
			//err2 = em.AddEndpoint(ctx2, key, endpoints.Endpoint{
			//	Addr:     addr,
			//	Metadata: now.String(),
			//})
			assert.NoError(t, err2)
		}
	}()

	// 3. 开启gRPC服务端
	server := grpc.NewServer()
	RegisterUserServiceServer(server, &Server{})
	t.Log("listening...")
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
	err = e.cli.Close()
	assert.NoError(t, err)

}

func (e *EtcdTestSuite) TestClient() {
	t := e.T()
	resolverEtcd, err := resolver.NewBuilder(e.cli)
	assert.NoError(t, err)

	// 一定要三个'/' !!!
	cc, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(resolverEtcd),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	userClient := NewUserServiceClient(cc)
	for {
		id := rand.Int63n(100)
		resp, err2 := userClient.GetById(context.Background(), &GetByIdRequest{
			Id: id,
		})

		assert.NoError(t, err2)

		t.Log("rpc ret:", resp.User)
		time.Sleep(time.Second * 2)
	}

}

func TestEtcd(t *testing.T) {
	suite.Run(t, &EtcdTestSuite{})
}
