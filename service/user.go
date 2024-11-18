package service

import (
	"fmt"
	"github.com/tittuvarghese/ss-go-core/crypto"
	"github.com/tittuvarghese/ss-go-core/jwt"
	"github.com/tittuvarghese/ss-go-customer-service/constants"
	"github.com/tittuvarghese/ss-go-customer-service/core/database"
	"github.com/tittuvarghese/ss-go-customer-service/models"
)

func CreateUser(user models.User, storage *database.RelationalDatabase) error {
	user.Password, _ = crypto.HashPassword(user.Password)
	err := storage.Instance.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

func AuthenticateUser(req models.LoginRequest, storage *database.RelationalDatabase) (string, error) {
	var user models.User
	condition := map[string]interface{}{"username": req.Username}

	// Pass a slice of User to QueryByCondition
	res, err := storage.Instance.QueryByCondition(&user, condition)
	if err != nil {
		return "", err
	}

	// Check if the result contains any user
	if len(res) == 0 {
		return "", fmt.Errorf("user not found")
	}

	// Cast the result to the correct type (since QueryByCondition returns []interface{})
	foundUser, ok := res[0].(*models.User)
	if !ok {
		return "", fmt.Errorf("type assertion failed")
	}

	err = crypto.ValidatePassword(foundUser.Password, req.Password)

	if err == nil {
		// Create JWT Token
		payload := models.AuthTokenPayload{
			Username:  foundUser.Username,
			ID:        foundUser.ID,
			Firstname: foundUser.Firstname,
			Lastname:  foundUser.Lastname,
			Type:      foundUser.Type,
		}

		token, err := jwt.Generate(payload, constants.AppName, constants.JwtExpiryHours)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	return "", fmt.Errorf("invalid password")
}

func GetProfile(userid string, storage *database.RelationalDatabase) (models.UserProfile, error) {
	var user models.User
	condition := map[string]interface{}{"id": userid}

	// Pass a slice of User to QueryByCondition
	res, err := storage.Instance.QueryByCondition(&user, condition)
	if err != nil {
		return models.UserProfile{}, err
	}

	// Check if the result contains any user
	if len(res) == 0 {
		return models.UserProfile{}, fmt.Errorf("user not found")
	}

	// Cast the result to the correct type (since QueryByCondition returns []interface{})
	foundUser, ok := res[0].(*models.User)
	if !ok {
		return models.UserProfile{}, fmt.Errorf("type assertion failed")
	}

	result := models.UserProfile{
		Username:  foundUser.Username,
		ID:        foundUser.ID,
		Firstname: foundUser.Firstname,
		Lastname:  foundUser.Lastname,
		Type:      foundUser.Type,
	}
	return result, nil

}
