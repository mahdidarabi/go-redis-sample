package models

import (
	"context"
	"go-redis-sample/internal/utils"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserProfile struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

type UserProfileDAO struct {
	redisClient         *redis.Client
	cacheExpirationTime time.Duration
}

var userProfileDAO = &UserProfileDAO{
	redisClient:         nil,
	cacheExpirationTime: 0,
}

var ctx = context.Background()

//#region Setters

func SetRedisClient(redisClient *redis.Client) *UserProfileDAO {
	userProfileDAO.redisClient = redisClient
	return userProfileDAO
}

func SetCacheExpirationTime(cacheExpirationTime time.Duration) *UserProfileDAO {
	userProfileDAO.cacheExpirationTime = cacheExpirationTime
	return userProfileDAO
}

//#endregion

//#region User Profile logics

func GetUserProfileById(id string) (UserProfile, error) {
	var userProfile UserProfile

	userProfileBytes, err := userProfileDAO.redisClient.Get(ctx, id).Bytes()
	if err != nil {
		log.Println("Error getting user profile:", err)
		return UserProfile{}, err
	}

	err = utils.UnmarshalJson(userProfileBytes, &userProfile)
	if err != nil {
		log.Println("Error getting user profile:", err)
		return UserProfile{}, err
	}
	return userProfile, nil
}

func SetUserProfile(id string, userProfile UserProfile) error {
	userProfileBytes, marshalErr := utils.MarshalJson(userProfile)
	if marshalErr != nil {
		log.Println("Error marshalling user profile:", marshalErr)
		return marshalErr
	}

	err := userProfileDAO.redisClient.Set(ctx, id, userProfileBytes, userProfileDAO.cacheExpirationTime).Err()
	if err != nil {
		log.Println("Error creating user profile:", err)
		return err
	}
	return nil
}

func DeleteUserProfile(userProfileId string) error {
	err := userProfileDAO.redisClient.Del(ctx, userProfileId).Err()
	if err != nil {
		return err
	}
	return nil
}

//#endregion
