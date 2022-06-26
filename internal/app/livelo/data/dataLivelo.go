package data

import (
	"context"
	"fmt"

	"github.com/pontuando/scraper-livelo/internal/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConfig struct {
	mongoURL string
	database string
	collection string
}

func NewMongoConfig(mongoURL string) *mongoConfig {
	return &mongoConfig{
		mongoURL: mongoURL,
		database: model.DatabaseName,
		collection: model.PartnerCollectionName,
	}
}


func (mc *mongoConfig) SavePartners(ctx context.Context, partners []model.PartnerTreated) error {
	mgCli, err := mongo.Connect(ctx, options.Client().ApplyURI(mc.mongoURL))
	if err != nil {
		return fmt.Errorf("fails attempting connect to mongodb: %s", err)
	}

	collection := mgCli.Database(mc.database).Collection(mc.collection)

	pi := make([]interface{}, len(partners))
	for i, v := range partners {
		pi[i] = v
	}

	_, err = collection.InsertMany(ctx, pi)
	if err != nil {
		return fmt.Errorf("fails attempting insert many: %s", err)
	}

	return nil
}

func (mc *mongoConfig) GetAllPartners(ctx context.Context) ([]model.PartnerTreated, error) {
	mgCli, err := mongo.Connect(ctx, options.Client().ApplyURI(mc.mongoURL))
	if err != nil {
		return nil, fmt.Errorf("fails attempting connect to mongodb: %s", err)
	}

	collection := mgCli.Database(mc.database).Collection(mc.collection)

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("fails attempting find partners: %s", err)
	}

	var partners []model.PartnerTreated
	if err = cur.All(ctx, &partners); err != nil {
		return nil, fmt.Errorf("fails attempting read partners: %s", err)
	}

	return partners, nil
}