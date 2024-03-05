package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"strings"
	"testing"
	"time"
)

type ElasticSearchStandardAPISuite struct {
	suite.Suite
	es *elastic.Client
}

func (e *ElasticSearchStandardAPISuite) SetupSuite() {
	es, err := elastic.NewClient(elastic.Config{
		Addresses: []string{esAddr},
	})
	assert.NoError(e.T(), err)
	e.es = es

	e.TestCreateIndex()

}

func (e *ElasticSearchStandardAPISuite) TearDownSuite() {
	t := e.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := e.es.Indices.Delete([]string{"user_idx_go"}, e.es.Indices.Delete.WithContext(ctx))
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func (e *ElasticSearchStandardAPISuite) TestCreateIndex() {
	t := e.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := e.es.Indices.Create("user_idx_go",
		e.es.Indices.Create.WithBody(strings.NewReader(createIndexBody)),
		e.es.Indices.Create.WithContext(ctx))
	assert.NoError(t, err)

}

func (e *ElasticSearchStandardAPISuite) TestPutDoc() {
	t := e.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	u := User{
		Name:     "test2",
		Email:    "222@qq.com",
		Phone:    "22222",
		Birthday: time.Now().UnixMilli(),
	}

	data, err := json.Marshal(&u)
	assert.NoError(t, err)

	_, err = e.es.Index("user_idx_go",
		bytes.NewReader(data),
		e.es.Index.WithContext(ctx))
	assert.NoError(t, err)
}

func (e *ElasticSearchStandardAPISuite) parseUser(str string) User {
	idx := strings.Index(str, "_source")
	if idx != -1 {
		res := str[idx:]
		st := strings.Index(res, "{")
		if st != -1 {
			ed := strings.Index(res, "}")
			var u User
			err := json.Unmarshal([]byte(res[st:ed+1]), &u)
			if err != nil {
				return User{}
			}
			return u
		}
	}

	return User{}
}

func (e *ElasticSearchStandardAPISuite) TestGetDoc() {
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

	_, err = e.es.Index("user_idx_go", bytes.NewReader(data), e.es.Index.WithContext(ctx))
	assert.NoError(t, err)

	query := `{
	"query": {
		"term": {
			"name": "test2"
		}
	}
}`

	resp, err := e.es.Search(e.es.Search.WithContext(ctx),
		e.es.Search.WithIndex("user_idx_go"),
		e.es.Search.WithBody(strings.NewReader(query)))

	assert.NoError(t, err)
	res, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	resUser := e.parseUser(string(res))
	assert.Equal(t, u, resUser)

}

func TestElasticSearchStandardAPI(t *testing.T) {
	suite.Run(t, &ElasticSearchStandardAPISuite{})
}
