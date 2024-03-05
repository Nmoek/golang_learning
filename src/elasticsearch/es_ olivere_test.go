package elasticsearch

import (
	"context"
	"encoding/json"
	olivere "github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ElasticSearchOlivereAPISuite struct {
	suite.Suite
	olivere *olivere.Client
}

func (e *ElasticSearchOlivereAPISuite) SetupSuite() {

	o, err := olivere.NewClient(olivere.SetURL(esAddr), olivere.SetSniff(false))
	assert.NoError(e.T(), err)
	e.olivere = o

	e.TestCreateIndex()
}

func (e *ElasticSearchOlivereAPISuite) TearDownSuite() {
	t := e.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := e.olivere.DeleteIndex("user_idx_go2").Do(ctx)
	assert.NoError(t, err)
}

func (e *ElasticSearchOlivereAPISuite) TestCreateIndex() {
	t := e.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := e.olivere.CreateIndex("user_idx_go2").Body(createIndexBody).Do(ctx)
	assert.NoError(t, err)
}

func (e *ElasticSearchOlivereAPISuite) TestPutDoc() {
	t := e.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	u := User{
		Name:     "test2",
		Email:    "222@qq.com",
		Phone:    "22222",
		Birthday: time.Now().UnixMilli(),
	}

	data, err := json.Marshal(&u)
	assert.NoError(t, err)

	_, err = e.olivere.Index().Index("user_idx_go2").BodyString(string(data)).Do(ctx)
	assert.NoError(t, err)

}

func (e *ElasticSearchOlivereAPISuite) TestGetDoc() {
	t := e.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	u := User{
		Name:     "test3",
		Email:    "3333@qq.com",
		Phone:    "33333",
		Birthday: time.Now().UnixMilli(),
	}

	data, err := json.Marshal(&u)
	assert.NoError(t, err)

	_, err = e.olivere.Index().Index("user_idx_go2").BodyString(string(data)).Do(ctx)
	assert.NoError(t, err)

	q := olivere.NewMatchQuery("name", "test3")
	res, err := e.olivere.Search("user_idx_go2").Query(q).Do(ctx)
	assert.NoError(t, err)

	var resUser User
	err = json.Unmarshal(res.Hits.Hits[0].Source, &resUser)
	assert.NoError(t, err)
	assert.Equal(t, u, resUser)

}

func TestElasticSearchOlivereAPI(t *testing.T) {
	suite.Run(t, &ElasticSearchOlivereAPISuite{})
}
