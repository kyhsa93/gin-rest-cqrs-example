package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface repository inteface
type Interface interface {
	Create(
		accountID string,
		email string,
		provider string,
		socialID string,
		password string,
	) (entity.Account, error)
	Update(
		accountID string,
		email string,
		provider string,
		socialID string,
		password string,
	) (entity.Account, error)
	FindByEmailAndProvider(
		email string, provider string, unscoped bool,
	) entity.Account
	FindByID(id string, unscoped bool) entity.Account
	Delete(id string) entity.Account
}

// Repository repository for query to database
type Repository struct {
	redis *redis.Client
	mongo *mongo.Collection
}

// New create repository instance
func New(redis *redis.Client, mongo *mongo.Collection) *Repository {
	return &Repository{mongo: mongo, redis: redis}
}

func (repository *Repository) setCache(key string, accountEntity *entity.Account) {
	marshaledEntity, _ := json.Marshal(&accountEntity)
	setRedisDataError := repository.redis.Set(
		"account:"+key, string(marshaledEntity), time.Second,
	).Err()
	if setRedisDataError != nil {
		log.Println("Set Data to Redis Error: ", setRedisDataError)
	}
}

func (repository *Repository) getCache(key string) *entity.Account {
	data, getDataFromRedisError := repository.redis.Get("account:" + key).Result()
	if getDataFromRedisError != nil {
		log.Println("Get Data from Redis Error: ", getDataFromRedisError)
		return nil
	}

	entity := &entity.Account{}
	jsonUnmarshalError := json.Unmarshal([]byte(data), entity)
	if jsonUnmarshalError != nil {
		log.Println("Fail to unmarshal cached data", jsonUnmarshalError)
		return nil
	}

	if entity.ID == "" {
		return nil
	}
	return entity
}

// Create create account
func (repository *Repository) Create(
	accountID string,
	email string,
	provider string,
	socialID string,
	password string,
) (entity.Account, error) {
	sameEmailAccount := entity.Account{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"email": email},
	).Decode(&sameEmailAccount)

	if sameEmailAccount.ID != "" {
		return sameEmailAccount, errors.New("Duplicated Email")
	}
	accountEntity := entity.Account{
		ID:        accountID,
		Email:     email,
		Provider:  provider,
		Password:  password,
		SocialID:  socialID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	insertResult, err := repository.mongo.InsertOne(context.TODO(), accountEntity)
	if err != nil || insertResult == nil {
		panic(err)
	}
	repository.setCache(accountID, &accountEntity)
	return accountEntity, nil
}

// Update update account
func (repository *Repository) Update(
	accountID string,
	email string,
	provider string,
	socialID string,
	password string,
) (entity.Account, error) {
	condition := bson.M{"_id": accountID}
	_, err := repository.mongo.UpdateOne(
		context.TODO(),
		condition,
		bson.M{
			"$set": bson.M{
				"updatedAt": time.Now(),
			},
		},
	)
	if err != nil {
		panic(err)
	}
	updated := entity.Account{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": accountID},
	).Decode(&updated)
	repository.setCache(accountID, &updated)
	return updated, nil
}

// FindByEmailAndProvider find all account
func (repository *Repository) FindByEmailAndProvider(
	email string,
	provider string,
	unscoped bool,
) entity.Account {
	accountEntity := entity.Account{}

	if unscoped == true {
		repository.mongo.FindOne(
			context.TODO(),
			bson.M{
				"email": email, "provider": provider,
				"$ne": []interface{}{bson.M{"deletedAt": nil}},
			},
		).Decode(accountEntity)
		return accountEntity
	}

	if cache := repository.getCache(email); cache != nil {
		return *cache
	}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"email": email, "provider": provider},
	).Decode(&accountEntity)
	repository.setCache(email, &accountEntity)

	return accountEntity
}

// FindByID find account by accountId
func (repository *Repository) FindByID(
	accountID string,
	unscoped bool,
) entity.Account {
	accountEntity := entity.Account{}

	if unscoped == true {
		repository.mongo.FindOne(
			context.TODO(),
			bson.M{
				"_id": accountID,
				"$ne": []interface{}{bson.M{"deletedAt": nil}},
			},
		).Decode(&accountEntity)
		return accountEntity
	}

	if cache := repository.getCache(accountID); cache != nil {
		return *cache
	}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": accountID, "deletedAt": nil},
	).Decode(&accountEntity)
	repository.setCache(accountID, &accountEntity)
	return accountEntity
}

// Delete delete account by accountId
func (repository *Repository) Delete(accountID string) entity.Account {
	accountEntity := entity.Account{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": accountID, "deletedAt": nil},
	).Decode(&accountEntity)
	condition := bson.M{"_id": accountID}
	repository.mongo.UpdateOne(
		context.TODO(),
		condition,
		bson.M{
			"$set": bson.M{
				"deletedAt": time.Now(),
			},
		},
	)
	return accountEntity
}
