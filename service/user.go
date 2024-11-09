package service

import (
	"github.com/tittuvarghese/customer-service/core/database"
	"github.com/tittuvarghese/customer-service/models"
)

func CreateUser(user models.User, storage *database.RelationalDatabase) error {
	err := storage.Instance.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}
