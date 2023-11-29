// Package mongodb
// @Description: mongodb操作测试
package mongodb_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

func TestMongoDB(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// 监视器
	monitor := &event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			fmt.Printf("[monitor msg]: %v \n", startedEvent.Command.String())
		},
	}
	opts := options.Client().
		ApplyURI("mongodb://localhost:27017").
		SetMonitor(monitor)

	client, err := mongo.Connect(ctx, opts)
	assert.NoError(t, err)

	err = client.Ping(ctx, readpref.Primary())
	assert.NoError(t, err)
	defer func() {
		err = client.Disconnect(ctx)
		assert.NoError(t, err)
	}()

	// 打开某个数据库中的某张表
	col := client.Database("kitbook").Collection("articles")

	// 插入文档
	insertRes, err := col.InsertOne(ctx, Article{
		Id:       1,
		Title:    "MongoDB标题",
		Content:  "MongoDB内容",
		AuthorId: 123,
		Status:   1,
	})
	assert.NoError(t, err)

	// 不指定InsertedID, 会自动生成primitive.ObjectID类型
	// 类似自增主键，但却是12byte, 需要与Mysql自增主键8byte做转换
	oid := insertRes.InsertedID.(primitive.ObjectID)
	t.Log("InsertedID: ", oid)

	// 查找文档
	//filterFind := bson.D{bson.E{"id", 1}}
	filterFind := bson.M{
		"id": 1,
	}
	findRes := col.FindOne(ctx, filterFind)
	if findRes.Err() == mongo.ErrNoDocuments {
		t.Error("没有找到数据")
	} else {
		assert.NoError(t, findRes.Err())

		var resArt Article
		err = findRes.Decode(&resArt)
		assert.NoError(t, err)

		t.Log(resArt)
	}

	// 更新文档
	filterUpdate := bson.M{
		"id": 1,
	}

	now := time.Now().UnixMilli()
	set := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				{"title", "Mongo新标题"},
				{"content", "Mongo新内容"},
				{"utime", now},
				{"ctime", now},
			},
		},
	}

	updateRes, err := col.UpdateOne(ctx, filterUpdate, set)
	assert.NoError(t, err)

	t.Log("update count:", updateRes.ModifiedCount)

	// 插入文档2
	insertRes, err = col.InsertOne(ctx, Article{
		Id:       2,
		Title:    "MongoDB标题",
		Content:  "MongoDB内容",
		AuthorId: 345,
		Status:   2,
	})

	filterFinds := bson.D{
		bson.E{"id", 1},
		bson.E{"id", 2},
	}

	findReses, err := col.Find(ctx, filterFinds)
	//assert.NoError(t, err)
	//
	count := findReses.RemainingBatchLength()
	t.Log("all count: ", count, ", raw: ", findReses.Current)

	//arts := make([]Article, count)
	//err = findReses.All(ctx, &arts)
	//assert.NoError(t, err)
	//
	//t.Log("all res: ", arts)

	// 删除文档
	filterDel := bson.D{
		bson.E{"id", 1},
	}
	delRes, err := col.DeleteOne(ctx, filterDel)
	assert.NoError(t, err)

	t.Log("delCount: ", delRes.DeletedCount)
}

type Article struct {
	Id       int64  `bson:"id,omitempty"`
	Title    string `bson:"title,omitempty"`
	Content  string `bson:"content,omitempty"`
	AuthorId int64  `bson:"author_id,omitempty"`
	Status   uint8  `bson:"status,omitempty"`
	Ctime    int64  `bson:"ctime,omitempty"`
	Utime    int64  `bson:"utime,omitempty"`
}
