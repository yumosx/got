package mongox

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Filter interface {
	bson.D | bson.M
}

type Builder[T Filter] struct {
	collection *mongo.Collection
	filter     T
	ctx        context.Context
}

func NewMongoBuilder[T Filter](collection *mongo.Collection) *Builder[T] {
	return &Builder[T]{collection: collection}
}

func (builder *Builder[T]) WithContext(ctx context.Context) *Builder[T] {
	builder.ctx = ctx
	return builder
}

func (builder *Builder[T]) Filter(filter T) *Builder[T] {
	builder.filter = filter
	return builder
}

func (builder *Builder[T]) InsertOne(value any) (*mongo.InsertOneResult, error) {
	if builder.ctx == nil {
		builder.ctx = context.TODO()
	}
	return builder.collection.InsertOne(builder.ctx, value)
}

func (builder *Builder[T]) FindOne() *mongo.SingleResult {
	return builder.collection.FindOne(builder.ctx, builder.filter)
}

func (builder *Builder[T]) UpdateOne(value bson.E) (*mongo.UpdateResult, error) {
	if builder.ctx == nil {
		builder.ctx = context.TODO()
	}

	d := bson.D{bson.E{Key: "$set", Value: value}}
	return builder.collection.UpdateOne(builder.ctx, builder.filter, d)
}

func (builder *Builder[T]) UpdateMany(value any) (*mongo.UpdateResult, error) {
	if builder.ctx == nil {
		builder.ctx = context.TODO()
	}

	d := bson.D{bson.E{Key: "$set", Value: value}}
	return builder.collection.UpdateMany(builder.ctx, builder.filter, d)
}

func (builder *Builder[T]) DeleteMany() (*mongo.DeleteResult, error) {
	if builder.ctx == nil {
		builder.ctx = context.TODO()
	}

	return builder.collection.DeleteMany(builder.ctx, builder.filter)
}
