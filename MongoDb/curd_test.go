package MongoDb

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMongoDB(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 设置monitor
	monitor := &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			fmt.Println(evt.Command)
		},
	}
	// 链接mongodb， 并设置监控
	client, err := mongo.Connect(
		ctx,
		options.
			Client().
			ApplyURI("mongodb://root:example@localhost:27017").
			SetMonitor(monitor),
	)
	assert.NoError(t, err)
	// 操作 client，设置数据库名和表名
	col := client.Database("small-book").Collection("articles")
	// 插入数据
	insertRes, err := col.InsertOne(
		ctx,
		Article{
			Id:       1,
			Title:    "我的标题",
			Content:  "我的内容",
			AuthorId: 123,
		},
	)
	// 断定应该不会有error
	assert.NoError(t, err)
	oid := insertRes.InsertedID.(primitive.ObjectID)
	t.Log("插入的id", oid.String())

	filter := bson.D{bson.E{Key: "id", Value: 1}}
	findRes := col.FindOne(ctx, filter)
	if findRes.Err() == mongo.ErrNoDocuments {
		t.Log("未找到数据")
	} else {
		assert.NoError(t, findRes.Err())

		var art Article
		err = findRes.Decode(&art)
		assert.NoError(t, err)
		t.Log(art)
	}

	// --- 更新
	// 设置更新条件
	updateFilter := bson.D{
		bson.E{
			Key:   "id",
			Value: 1,
		},
	}
	set := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.E{
				Key:   "title",
				Value: "要被更新的title",
			},
		},
	}
	updateOneRes, err := col.UpdateOne(ctx, updateFilter, set)
	// 断言不会发生错误
	assert.NoError(t, err)
	t.Log("更新文档数量: ", updateOneRes.MatchedCount)

	// 批量更新，
	updateManyRes, err := col.UpdateMany(
		ctx,
		updateFilter,
		bson.D{
			bson.E{
				Key: "$set",
				// 这里会把article整个结构体序列化更新，由于没有忽略零值， 未设置的字段会被默认为该类型的零值， 这里添加一个示例
				// 设置id忽略零值
				Value: Article{Content: "新的内容"},
			},
		},
	)
	// 断言没有错误
	assert.NoError(t, err)
	t.Log("更新文档数量: ", updateManyRes.ModifiedCount)

	// 删除文档
	deleteFilter := bson.D{
		bson.E{
			Key:   "id",
			Value: 1,
		},
	}
	delRes, err := col.DeleteMany(ctx, deleteFilter)
	assert.NoError(t, err)
	t.Log("删除文档数量：", delRes.DeletedCount)

}

type Article struct {
	Id       int64 `bson:"id,omitempty"`
	Title    string
	Content  string
	AuthorId int64
	Status   uint8
	Ctime    int64
	Utime    int64
}
