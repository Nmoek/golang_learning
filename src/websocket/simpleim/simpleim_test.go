package simpleim

import (
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GatewayTestSuite struct {
	suite.Suite
	client sarama.Client
}

func (g *GatewayTestSuite) SetupSuite() {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true

	client, err := sarama.NewClient([]string{"localhost:9094"}, cfg)
	assert.NoError(g.T(), err)
	g.client = client
}

func (g *GatewayTestSuite) TestGateway() {

	t := g.T()
	go func() {
		err := g.startGateway(":8081", "client_8081")
		t.Log("8081退出服务", err)
	}()

	go func() {
		err := g.startGateway(":8082", "client_8082")
		t.Log("8082退出服务", err)
	}()

	err := g.startGateway(":8083", "client_8083")
	t.Log("8083退出服务", err)

}

func (g *GatewayTestSuite) startGateway(addr string, instanceId string) error {
	producer, err := sarama.NewSyncProducerFromClient(g.client)
	if err != nil {
		return err
	}

	ws := NewWsGateway(
		&Service{
			producer: producer,
		}, g.client, instanceId)

	return ws.Start(addr)
}

func TestWsGateway(t *testing.T) {
	suite.Run(t, &GatewayTestSuite{})
}
