package repository

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface repository interface
type Interface interface {
	Create(
		profileID string,
		accountID string,
		email string,
		gender string,
		fileID string,
		interestedField string,
		interestedFieldDetail []string,
	) (entity.Profile, error)
	FindByID(
		profileID string,
	) (entity.Profile, error)
	FindByAccountID(
		accountID string,
	) (entity.Profile, error)
}

// Repository repository for profile data
type Repository struct {
	redis *redis.Client
	mongo *mongo.Collection
}

// New create repository instance
func New(redis *redis.Client, mongo *mongo.Collection) *Repository {
	return &Repository{mongo: mongo, redis: redis}
}

func (repository *Repository) setCache(key string, profileEntity *entity.Profile) {
	marshaledEntity, _ := json.Marshal(&profileEntity)
	setRedisDataError := repository.redis.Set(
		"profile:"+key, string(marshaledEntity), time.Second,
	).Err()
	if setRedisDataError != nil {
		log.Println("Set Data to Redis Error: ", setRedisDataError)
	}
}

func (repository *Repository) getCache(key string) *entity.Profile {
	data, getDataFromRedisError := repository.redis.Get("profile:" + key).Result()
	if getDataFromRedisError != nil {
		log.Println("Get Data from Redis Error: ", getDataFromRedisError)
		return nil
	}

	entity := &entity.Profile{}
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

// Create create profile data in databse
func (repository *Repository) Create(
	profileID string,
	accountID string,
	email string,
	gender string,
	fileID string,
	interestedField string,
	interestedFieldDetail []string,
) (entity.Profile, error) {
	profileEntity := entity.Profile{
		ID:                    profileID,
		AccountID:             accountID,
		Email:                 email,
		Gender:                gender,
		FileID:                fileID,
		InterestedField:       interestedField,
		InterestedFieldDetail: interestedFieldDetail,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
	insertResult, err := repository.mongo.InsertOne(
		context.TODO(),
		profileEntity,
	)
	if err != nil || insertResult == nil {
		panic(err)
	}
	repository.setCache(profileID, &profileEntity)
	return profileEntity, nil
}

// FindByID find profile data by profile id
func (repository *Repository) FindByID(
	profileID string,
) (entity.Profile, error) {
	profileEntity := entity.Profile{}
	if cache := repository.getCache(profileID); cache != nil {
		return *cache, nil
	}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": profileID, "deletedAt": nil},
	).Decode(&profileEntity)
	return profileEntity, nil
}

// FindByAccountID find profile by accountID
func (repository *Repository) FindByAccountID(
	accountID string,
) (entity.Profile, error) {
	profileEntity := entity.Profile{}
	if cache := repository.getCache(accountID); cache != nil {
		return *cache, nil
	}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"accountId": accountID, "deletedAt": nil},
	).Decode(&profileEntity)
	return profileEntity, nil
}
