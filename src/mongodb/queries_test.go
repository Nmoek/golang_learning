package mongodb_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

type MongoDBTestSuite struct {
	suite.Suite
	col *mongo.Collection
}

func (m *MongoDBTestSuite) TearDownTest() {
	t := m.T()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// 清空所有数据
	delRes, err := m.col.DeleteMany(ctx, bson.D{})
	assert.NoError(t, err)

	// 把所有索引去掉
	_, err = m.col.Indexes().DropAll(ctx)
	assert.NoError(t, err)

	t.Log("del count: ", delRes.DeletedCount)
}

func (m *MongoDBTestSuite) SetupSuite() {
	t := m.T()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	monitor := &event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			fmt.Printf("[mongo msg]: %v \n", startedEvent.Command)
		},
	}

	opts := options.Client().
		ApplyURI("mongodb://localhost:27017").
		SetMonitor(monitor)

	client, err := mongo.Connect(ctx, opts)
	err = client.Ping(ctx, readpref.Primary())
	assert.NoError(t, err)

	col := client.Database("kitbook").Collection("articles")

	m.col = col
	now := time.Now().UnixMilli()

	manyRes, err := col.InsertMany(ctx, []any{Article{
		Id:       1,
		Title:    "Mongo测试套件标题",
		Content:  "Mongo测试套件内容",
		AuthorId: 123,
		Ctime:    now,
		Utime:    now,
	}, Article{
		Id:       2,
		Title:    "Mongo测试套件标题",
		Content:  "Mongo测试套件内容",
		AuthorId: 456,
		Ctime:    now,
		Utime:    now,
	}})

	assert.NoError(t, err)

	t.Log("count: ", len(manyRes.InsertedIDs), "InsertIDs: ", manyRes.InsertedIDs)

}

// @func: TestOr
// @date: 2023-11-29 22:09:18
// @brief: mongodb复杂查询-Or查询
// @author: Kewin Li
// @receiver m
func (m *MongoDBTestSuite) TestOr() {
	t := m.T()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	filterFind := bson.A{
		bson.D{bson.E{"id", 1}},
		bson.D{bson.E{"author_id", 123}},
	}

	findRes, err := m.col.Find(ctx, bson.D{bson.E{"$or", filterFind}})
	assert.NoError(t, err)

	var arts []Article
	err = findRes.All(ctx, &arts)
	assert.NoError(t, err)

	t.Log("all res: ", arts)
}

// @func: TestAnd
// @date: 2023-11-30 02:28:43
// @brief: mongodb复杂查询-And
// @author: Kewin Li
// @receiver m
func (m *MongoDBTestSuite) TestAnd() {
	t := m.T()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	filter := bson.A{
		bson.D{bson.E{"id", 2}},
		bson.D{bson.E{"author_id", 123}},
	}

	findRes, err := m.col.Find(ctx, bson.D{bson.E{"$and", filter}})
	assert.NoError(t, err)

	var arts []Article
	err = findRes.All(ctx, &arts)
	assert.NoError(t, err)

	t.Log("count: ", len(arts), ", res: ", arts)
}

// @func: TestIn
// @date: 2023-11-30 03:15:54
// @brief: mongodb复杂查询-In
// @author: Kewin Li
// @receiver m
func (m *MongoDBTestSuite) TestIn() {
	t := m.T()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	findRes, err := m.col.Find(ctx, bson.D{
		bson.E{
			"id",
			bson.D{bson.E{"$in", []int{1, 456}}},
		},
	})

	assert.NoError(t, err)

	var arts []Article
	err = findRes.All(ctx, &arts)
	assert.NoError(t, err)

	t.Log("count: ", len(arts), "all res: ", arts)

}

// @func: TestFindAll
// @date: 2023-11-30 03:27:40
// @brief: 全表查询
// @author: Kewin Li
// @receiver m
func (m *MongoDBTestSuite) TestFindAll() {
	t := m.T()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	findRes, err := m.col.Find(ctx, bson.D{})
	assert.NoError(t, err)

	var arts []Article

	err = findRes.All(ctx, &arts)
	assert.NoError(t, err)

	t.Log("all res", arts)
}

// @func: TestProjection
// @date: 2023-11-30 03:40:08
// @brief: 部分字段查询
// @author: Kewin Li
// @receiver m
func (m *MongoDBTestSuite) TestProjection() {
	t := m.T()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//proj1 := bson.D{
	//	{"id", 1},
	//	//{"author_id", 123},
	//}

	proj2 := bson.M{"id": 1, "author_id": 1}
	findRes, err := m.col.Find(ctx, bson.D{},
		options.Find().SetProjection(proj2))
	assert.NoError(t, err)

	var arts []Article
	err = findRes.All(ctx, &arts)
	assert.NoError(t, err)

	t.Log("count:", len(arts), arts)

}

// @func: TestIndexes
// @date: 2023-11-30 03:40:30
// @brief: 创建索引
// @author: Kewin Li
// @receiver m
func (m *MongoDBTestSuite) TestIndexes() {
	t := m.T()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	str, err := m.col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{"id", 1},
			{"author_id", 1},
		},
		Options: options.Index().SetUnique(true), //设置为唯一索引
	})

	strs, err := m.col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{
				"id": 1,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{
				"author_id": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	})

	assert.NoError(t, err)

	t.Log("indexes:", strs)
}

func TestMongoDBQueries(t *testing.T) {
	suite.Run(t, &MongoDBTestSuite{})
}
