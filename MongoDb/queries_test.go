package MongoDb

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

type MongoDBTestSuite struct {
	suite.Suite
	col *mongo.Collection
}

func (s *MongoDBTestSuite) SetupSuite() {
	t := s.T()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 设置monitor
	monitor := &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			fmt.Println(evt.Command)
		},
	}
	opts := options.
		Client().
		ApplyURI("mongodb://root:example@localhost:27017").
		SetMonitor(monitor)
	// 链接mongodb， 并设置监控
	client, err := mongo.Connect(
		ctx,
		opts,
	)
	assert.NoError(t, err)
	// 操作 client，设置数据库名和表名
	col := client.
		Database("small-book").
		Collection("articles")
	s.col = col

	manyRes, err := col.InsertMany(ctx, []any{
		Article{
			Id:      123,
			Title:   "My Title",
			Content: "My Content",
		},
		Article{
			Id:      234,
			Title:   "My Title 1",
			Content: "My Content 1",
		},
	})
	assert.NoError(s.T(), err)
	s.T().Log("插入数量:", len(manyRes.InsertedIDs))
}

func (s *MongoDBTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	deleteRes, err := s.col.DeleteMany(ctx, bson.D{})
	assert.NoError(s.T(), err)
	s.T().Log("删除的条数：", deleteRes.DeletedCount)

	// 删除索引
	_, err = s.col.Indexes().DropAll(ctx)
	assert.NoError(s.T(), err)
}

func (s *MongoDBTestSuite) TestOr() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	filter := bson.A{
		bson.D{
			bson.E{
				Key:   "id",
				Value: 123,
			},
		},
		bson.D{
			bson.E{
				Key:   "id",
				Value: 234,
			},
		},
	}
	findRes, err := s.col.Find(
		ctx,
		bson.D{
			bson.E{
				Key:   "$or",
				Value: filter,
			},
		},
	)
	assert.NoError(s.T(), err)
	var arts []Article
	err = findRes.All(ctx, &arts)
	s.T().Log("查询结构：", arts)
}

func (s *MongoDBTestSuite) TestAnd() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	filter := bson.A{
		bson.D{
			bson.E{
				Key:   "id",
				Value: 123,
			},
		},
		bson.D{
			bson.E{
				Key:   "authorid",
				Value: 0,
			},
		},
	}
	findRes, err := s.col.Find(
		ctx,
		bson.D{
			bson.E{
				Key:   "$and",
				Value: filter,
			},
		},
	)
	assert.NoError(s.T(), err)
	var arts []Article
	err = findRes.All(ctx, &arts)
	s.T().Log("查询结构：", arts)
}

func (s *MongoDBTestSuite) TestIn() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	filter := bson.D{
		bson.E{
			Key: "id",
			Value: bson.D{
				bson.E{Key: "$in", Value: []int{123, 235}},
			},
		},
	}
	findRes, err := s.col.Find(
		ctx,
		filter,
		// 设置查询字段,只查询部分字段
		options.Find().SetProjection(
			bson.D{
				bson.E{Key: "id", Value: 1},
				bson.E{Key: "content", Value: 1},
				bson.E{Key: "title", Value: 1},
			},
		),
	)
	assert.NoError(s.T(), err)
	var arts []Article
	err = findRes.All(ctx, &arts)
	s.T().Log("查询结构：", arts)
}

func (s *MongoDBTestSuite) TestIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	iRes, err := s.col.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.D{bson.E{Key: "id", Value: 1}},
			// 自动设置索引名称
			//Options: options.Index().SetUnique(true),
			// 手动设置索引名称
			Options: options.Index().SetName("cus_index"),
		},
	)
	assert.NoError(s.T(), err)
	s.T().Log("创建索引：", iRes)

}

func TestMongoDBQueries(t *testing.T) {
	suite.Run(t, &MongoDBTestSuite{})

}
