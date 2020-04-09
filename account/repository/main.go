package repository

import (
	"context"
	"encoding/json"
	"errors"
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
		fcmToken string,
		gender string,
	) (entity.Account, error)
	Update(
		accountID string,
		password string,
		fcmToken string,
	) (entity.Account, error)
	FindByEmailAndProvider(
		email string, provider string, deleted bool,
	) entity.Account
	FindByEmail(email string, deleted bool) entity.Account
	FindByID(id string, deleted bool) entity.Account
	Delete(id string) entity.Account
}

// Repository repository for query to database
type Repository struct {
	redis *redis.Client
	mongo *mongo.Collection
}

// New create repository instance
func New(
	redis *redis.Client, mongo *mongo.Collection,
) Interface {
	return &Repository{mongo: mongo, redis: redis}
}

func (repository *Repository) setCache(
	key string, accountEntity *entity.Account,
) {
	marshaledEntity, _ := json.Marshal(&accountEntity)
	repository.redis.Set(
		"account:"+key, string(marshaledEntity), time.Second,
	)
}

func (repository *Repository) getCache(
	key string,
) *entity.Account {
	data, getDataFromRedisError :=
		repository.redis.Get("account:" + key).Result()
	if data == "" || getDataFromRedisError != nil {
		return nil
	}

	entity := &entity.Account{}
	jsonUnmarshalError := json.Unmarshal([]byte(data), entity)
	if entity.ID == "" || jsonUnmarshalError != nil {
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
	fcmToken string,
	gender string,
) (entity.Account, error) {
	sameEmailAccount := entity.Account{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"email": email, "deletedAt": nil},
	).Decode(&sameEmailAccount)

	if sameEmailAccount.ID != "" {
		return sameEmailAccount, errors.New("duplicated email")
	}
	accountEntity := entity.Account{
		ID:        accountID,
		Email:     email,
		Provider:  provider,
		Password:  password,
		SocialID:  socialID,
		FCMToken:  fcmToken,
		Gender:    gender,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	insertResult, err := repository.mongo.InsertOne(
		context.TODO(),
		accountEntity,
	)
	if err != nil || insertResult == nil {
		panic(err)
	}
	repository.setCache(accountID, &accountEntity)
	return accountEntity, nil
}

// Update update account
func (repository *Repository) Update(
	accountID string,
	password string,
	fcmToken string,
) (entity.Account, error) {
	account := entity.Account{}
	condition := bson.M{"_id": accountID, "deletedAt": nil}
	repository.mongo.FindOne(
		context.TODO(),
		condition,
	).Decode(&account)
	if account.ID == "" {
		return account, errors.New("update targe not found")
	}
	updateResult, err := repository.mongo.UpdateOne(
		context.TODO(),
		condition,
		bson.M{
			"$set": bson.M{
				"password":  password,
				"fcmToken":  fcmToken,
				"updatedAt": time.Now(),
			},
		},
	)
	if updateResult == nil || err != nil {
		panic(err)
	}
	repository.setCache(accountID, &account)
	return account, nil
}

// FindByEmailAndProvider find all account
func (repository *Repository) FindByEmailAndProvider(
	email string,
	provider string,
	deleted bool,
) entity.Account {
	accountEntity := entity.Account{}

	if cache := repository.getCache(email); cache != nil {
		return *cache
	}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{
			"email":     email,
			"provider":  provider,
			"deletedAt": nil,
		},
	).Decode(&accountEntity)
	repository.setCache(email, &accountEntity)
	return accountEntity
}

// FindByEmail find account by email
func (repository *Repository) FindByEmail(
	email string,
	deleted bool,
) entity.Account {
	accountEntity := entity.Account{}

	if deleted == true {
		repository.mongo.FindOne(
			context.TODO(),
			bson.M{"email": email},
		).Decode(&accountEntity)
		return accountEntity
	}

	if cache := repository.getCache(email); cache != nil {
		return *cache
	}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"email": email, "deletedAt": nil},
	).Decode(&accountEntity)
	repository.setCache(email, &accountEntity)
	return accountEntity
}

// FindByID find account by accountId
func (repository *Repository) FindByID(
	accountID string,
	deleted bool,
) entity.Account {
	accountEntity := entity.Account{}

	if deleted == true {
		repository.mongo.FindOne(
			context.TODO(),
			bson.M{
				"_id": accountID,
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
func (repository *Repository) Delete(
	accountID string,
) entity.Account {
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
