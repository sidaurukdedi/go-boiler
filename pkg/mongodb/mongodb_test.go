package mongodb_test

import (
	"context"
	"testing"

	"github.com/sidaurukdedi/go-boiler/pkg/mongodb"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client mongodb.Client
)

func TestMain(m *testing.M) {
	mockClientOptions := options.ClientOptions{}
	client = mongodb.NewClientAdapter(&mockClientOptions)
	m.Run()
}

func TestClientAdapter_Connect(t *testing.T) {
	err := client.Connect(context.TODO())
	assert.NoError(t, err)
}

func TestClientAdapter_Disconnect(t *testing.T) {
	err := client.Disconnect(context.TODO())
	assert.NoError(t, err)
}

func TestClientAdapter_Database(t *testing.T) {
	err := client.Connect(context.TODO())
	assert.NoError(t, err)
}

func TestDatabaseAdapter_Collection(t *testing.T) {
	collection := client.Database("test-db").Collection("test-collection")
	assert.NotNil(t, collection)
}

func TestCollectionAdapter_FindOne(t *testing.T) {
	result := client.Database("test-db").Collection("test-collection").FindOne(context.TODO(), make(map[string]interface{}))
	assert.NotNil(t, result)
	assert.Error(t, result.Err())
}

func TestCollectionAdapter_Find(t *testing.T) {
	cursor, err := client.Database("test-db").Collection("test-collection").Find(context.TODO(), make(map[string]interface{}))
	assert.Error(t, err)
	assert.Nil(t, cursor)
}

func TestCollectionAdapter_InsertOne(t *testing.T) {
	doc := make(map[string]interface{})
	result, err := client.Database("test-db").Collection("test-collection").InsertOne(context.TODO(), doc)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCollectionAdapter_InsertMany(t *testing.T) {
	docs := make([]interface{}, 0)
	doc := make(map[string]interface{})

	docs = append(docs, doc)

	result, err := client.Database("test-db").Collection("test-collection").InsertMany(context.TODO(), docs)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCollectionAdapter_CountDocuments(t *testing.T) {
	counted, err := client.Database("test-db").Collection("test-collection").CountDocuments(context.TODO(), make(map[string]interface{}))
	assert.Error(t, err)
	assert.Equal(t, int64(0), counted)
}

func TestCollectionAdapter_DeleteOne(t *testing.T) {
	result, err := client.Database("test-db").Collection("test-collection").DeleteOne(context.TODO(), make(map[string]interface{}))
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCollectionAdapter_DeleteMany(t *testing.T) {
	result, err := client.Database("test-db").Collection("test-collection").DeleteMany(context.TODO(), make(map[string]interface{}))
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCollectionAdapter_UpdateOne(t *testing.T) {
	updatedDoc := make(map[string]interface{})
	filter := make(map[string]interface{})

	result, err := client.Database("test-db").Collection("test-collection").UpdateOne(context.TODO(), filter, updatedDoc)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCollectionAdapter_UpdateMany(t *testing.T) {
	updatedDoc := make(map[string]interface{})
	filter := make(map[string]interface{})

	result, err := client.Database("test-db").Collection("test-collection").UpdateMany(context.TODO(), filter, updatedDoc)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCollectionAdapter_BulkWrite(t *testing.T) {
	bunchOfWriteModels := make([]mongo.WriteModel, 0)
	writeModel := mongo.NewInsertOneModel()

	bunchOfWriteModels = append(bunchOfWriteModels, writeModel)

	result, err := client.Database("test-db").Collection("test-collection").BulkWrite(context.TODO(), bunchOfWriteModels)
	assert.Error(t, err)
	assert.NotNil(t, result)
}
