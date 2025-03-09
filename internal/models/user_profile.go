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

type UserProfileModel struct {
	redisClient         *redis.Client
	cacheExpirationTime time.Duration
}

var userProfileModel = &UserProfileModel{
	redisClient:         nil,
	cacheExpirationTime: 0,
}

var ctx = context.Background()

//#region Setters

func SetRedisClient(redisClient *redis.Client) *UserProfileModel {
	userProfileModel.redisClient = redisClient
	return userProfileModel
}

func SetCacheExpirationTime(cacheExpirationTime time.Duration) *UserProfileModel {
	userProfileModel.cacheExpirationTime = cacheExpirationTime
	return userProfileModel
}

//#endregion

//#region User Profile logics

func SetUserProfile(id string, userProfile UserProfile) error {
	// if userProfile.Id == "" {
	// 	return errors.New("user profile ID is required")
	// }

	userProfileBytes, marshalErr := utils.MarshalJson(userProfile)
	if marshalErr != nil {
		log.Println("Error marshalling user profile:", marshalErr)
		return marshalErr
	}

	err := userProfileModel.redisClient.Set(ctx, id, userProfileBytes, userProfileModel.cacheExpirationTime).Err()
	if err != nil {
		log.Println("Error creating user profile:", err)
		return err
	}
	return nil
}

func GetUserProfileById(id string) (UserProfile, error) {
	var userProfile UserProfile

	userProfileBytes, err := userProfileModel.redisClient.Get(ctx, id).Bytes()
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

// func UpdateUserProfile(id string, userProfile UserProfile) error {
// 	err := userProfileModel.redisClient.Set(ctx, id, userProfile, userProfileModel.cacheExpirationTime).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func DeleteUserProfile(userProfileId string) error {
	err := userProfileModel.redisClient.Del(ctx, userProfileId).Err()
	if err != nil {
		return err
	}
	return nil
}

//#endregion
